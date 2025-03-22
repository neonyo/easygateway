package gw

import (
	"github.com/go-chi/chi/v5"
	"github.com/neonyo/gw/pkg/util"
	"github.com/neonyo/gw/router"
	"github.com/opentracing/opentracing-go"
	"net/http"
)

type httpProxy struct {
	conf     HttpConfig
	Span     opentracing.Span
	mux      util.SyncMap[string, *chi.Mux]
	endpoint util.SyncMap[string, *router.Endpoint]
}

func (g *httpProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if mux, ok := g.mux.Load(g.conf.hostSwitch(r)); ok {
		mux.ServeHTTP(w, r)
		return
	}
}
