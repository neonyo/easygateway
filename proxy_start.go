package easygateway

import (
	"errors"
	"fmt"
	"github.com/soheilhy/cmux"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"log"
	"net"
	"net/http"
)

func (p *proxy) start() error {
	if p.httpProxy.endpointOption == nil {
		return errors.New("需要设置选择服务站点方法 WithEndpointOption")
	}

	p.startHTTP()
	return nil
}

func (p *proxy) Stop() {

}

func (p *proxy) startHTTPWithListener(l net.Listener) {
	s := p.newHTTPServer()
	err := s.Serve(l)
	if err != nil {
		log.Fatalf("start http listeners failed with %+v", err)
	}
}

func (p *proxy) startHTTP() {
	l, err := net.Listen("tcp", p.cfg.Addr)
	if err != nil {
		log.Fatalf("start http failed failed with %+v", err)
	}
	m := cmux.New(l)
	go p.startHTTPWithListener(m.Match(cmux.Any()))
	fmt.Printf("Starting http server at %s...\n", p.cfg.Addr)
	err = m.Serve()
	if err != nil {
		log.Fatalf("start http failed failed with %+v", err)
	}
}

func (p *proxy) newHTTPServer() *http.Server {
	if !p.cfg.Telemetry {
		return &http.Server{
			Handler: p.httpProxy,
		}
	}
	return &http.Server{
		Handler: otelhttp.NewHandler(p.httpProxy, "otel-go-tracer", p.mwOptions()...),
	}
}

func (p *proxy) mwOptions() []otelhttp.Option {
	var options []otelhttp.Option
	options = append(options, otelhttp.WithSpanNameFormatter(func(operation string, r *http.Request) string {
		return fmt.Sprintf("%s", r.URL.Path)
	}))
	return options
}
