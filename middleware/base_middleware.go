package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/neonyo/easygateway/router"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"time"
)

type Base struct {
	Endpoint *router.Endpoint
	Router   *router.Router
}

type Baser interface {
	//GetHandlerType() int8
	GetEndpoint() *router.Endpoint
	GetRouter() *router.Router
	//GetLmt() *limiter.Limiter
	//GetMt() *core.Metrics
	//SetHandlerType(int8)
	SetEndpoint(*router.Endpoint)
	SetRouter(*router.Router)
	//SetLmt(*limiter.Limiter)
	//SetMt(*core.Metrics)
}

func (mw *Base) Init() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gwRw := &gatewayResponse{}
			gwRw.w = w
			gwRw.startTime = time.Now()
			gwRw.pErrorChan = make(chan error, 1)
			gwRw.pSuccessChan = make(chan bool, 1)
			reqBytes, _ := io.ReadAll(r.Body)
			buf := bytes.NewBuffer(reqBytes)
			r.Body = io.NopCloser(buf)
			ctx := r.Context()
			span := trace.SpanFromContext(ctx)
			if span.IsRecording() {
				headerJson, _ := json.Marshal(r.Header)
				span.SetAttributes(attribute.String("http.request.header", string(headerJson)))
				span.SetAttributes(attribute.String("http.request.body", string(reqBytes)))
			}
			w.Header().Set("Content-Type", "application/json")
			next.ServeHTTP(gwRw, r)
		})
	}
}

func (mw *Base) GetEndpoint() *router.Endpoint {
	return mw.Endpoint
}

func (mw *Base) GetRouter() *router.Router {
	return mw.Router
}

//func (mw *Base) GetLmt() *limiter.Limiter {
//	return mw.lmt
//}

//func (mw *Base) GetDomainLmt() *limiter.Limiter {
//	return mw.domainLmt
//}

//func (mw *Base) GetMt() *core.Metrics {
//	return mw.mt
//}

func (mw *Base) SetEndpoint(domain *router.Endpoint) {
	mw.Endpoint = domain
}

func (mw *Base) SetRouter(p *router.Router) {
	mw.Router = p
}

//func (mw *Base) SetLmt(lmt *limiter.Limiter) {
//	mw.lmt = lmt
//}

//func (mw *Base) SetDomainLmt(lmt *limiter.Limiter) {
//	mw.domainLmt = lmt
//}

//func (mw *Base) SetMt(mt *core.Metrics) {
//	mw.mt = mt
//}

//func (mw *Base) SetTraceId(traceId string) {
//	mw.traceId = traceId
//}

//func (mw *Base) GetTraceId() string {
//	return mw.traceId
//}
