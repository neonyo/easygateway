package easygateway

import (
	"github.com/go-chi/chi/v5"
	"github.com/neonyo/easygateway/pkg/util"
	"github.com/neonyo/easygateway/router"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type httpProxy struct {
	Telemetry      bool
	Span           opentracing.Span
	endpointOption endpointOption
	mux            util.SyncMap[string, *chi.Mux]
	endpoint       util.SyncMap[string, *router.Endpoint]
}

func (g *httpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mux, ok := g.mux.Load(g.endpointOption(r)); ok {
		mux.ServeHTTP(w, r)
		return
	}

}
