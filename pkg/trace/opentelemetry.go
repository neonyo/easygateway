package trace

import (
	"context"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.15.0"
	"log"
	"os"
)

func InitOpenTelemetry(serviceName, httpEndpoint, httpUrlPath string) (*otlptrace.Exporter, context.Context) {
	ctx := context.Background()

	// 创建一个将追踪信息输出到标准输出的导出器
	//exp, _ := stdouttrace.New(stdouttrace.WithPrettyPrint())

	var traceExporter *otlptrace.Exporter
	var batchSpanProcessor sdktrace.SpanProcessor

	traceExporter, batchSpanProcessor = newHTTPExporterAndSpanProcessor(ctx, httpEndpoint, httpUrlPath)

	otelResource := newResource(ctx, serviceName)

	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(otelResource),
		sdktrace.WithSpanProcessor(batchSpanProcessor),
		//sdktrace.WithBatcher(exp),
	)

	otel.SetTracerProvider(traceProvider)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return traceExporter, ctx
}

func newHTTPExporterAndSpanProcessor(ctx context.Context, httpEndpoint, httpUrlPath string) (*otlptrace.Exporter, sdktrace.SpanProcessor) {

	traceExporter, err := otlptrace.New(ctx, otlptracehttp.NewClient(
		otlptracehttp.WithEndpoint(httpEndpoint),
		otlptracehttp.WithURLPath(httpUrlPath),
		otlptracehttp.WithInsecure(),
		otlptracehttp.WithCompression(1)))

	if err != nil {
		log.Fatalf("%s: %v", "Failed to create the OpenTelemetry trace exporter", err)
	}

	batchSpanProcessor := sdktrace.NewBatchSpanProcessor(traceExporter)

	return traceExporter, batchSpanProcessor
}

func newResource(ctx context.Context, serviceName string) *resource.Resource {
	// hostname默认值为本机主机名
	hostName, _ := os.Hostname()
	r, err := resource.New(
		ctx,
		resource.WithFromEnv(),
		resource.WithProcess(), // runtime信息 process.runtime.name: go/gc, process.runtime.version: go1.20.1s
		resource.WithTelemetrySDK(),
		resource.WithHost(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String(serviceName),
			semconv.HostNameKey.String(hostName),
		),
	)
	rs, err := resource.Merge(resource.Default(), r)

	if err != nil {
		log.Fatalf("%s: %v", "Failed to create OpenTelemetry resource", err)
	}
	return rs
}
