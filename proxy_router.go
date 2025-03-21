package easygateway

import (
	"github.com/go-chi/chi/v5"
	"github.com/neonyo/easygateway/middleware"
	"github.com/neonyo/easygateway/router"
	"net/http"
	"strings"
)

// refresh 是否强制重新实例化
func (g *httpProxy) newMux(addr string, refresh bool) *chi.Mux {
	var mux = new(chi.Mux)
	var ok bool
	if mux, ok = g.mux.Load(addr); !ok || refresh {
		mux = chi.NewMux()
		g.mux.Store(addr, mux)
	}
	return mux
}

// bindEndpoint 添加反向代理服务路由
func (g *httpProxy) addEndpoints(rs []*router.Endpoint) {
	for _, e := range rs {
		g.endpoint.Store(e.Addr, e)
		g.muxEndpointHandle(g.newMux(e.Addr, false), e)
	}
}

func (g *httpProxy) updateEndpoint(e *router.Endpoint) {
	g.endpoint.Store(e.Addr, e)
	g.muxEndpointHandle(g.newMux(e.Addr, true), e)
}

func (g *httpProxy) delEndpoint(e *router.Endpoint) {
	g.mux.Delete(e.Addr)
	g.endpoint.Delete(e.Addr)
}

// 添加代理路由
func (g *httpProxy) addRouter(e *router.Endpoint, rs []*router.Router) {
	g.endpoint.Store(e.Addr, e)
	for _, v := range rs {
		g.muxRouterHandle(g.newMux(e.Addr, false), e, v)
	}
}

func (g *httpProxy) updateRouter(e *router.Endpoint) {
	g.endpoint.Store(e.Addr, e)
	g.muxEndpointHandle(g.newMux(e.Addr, true), e)
}

func (g *httpProxy) muxEndpointHandle(mux *chi.Mux, e *router.Endpoint) {
	if len(e.Router) > 0 {
		for _, r := range e.Router {
			g.mapMethodHandler(mux, r.ReqMethod, r.ReqPath, middleware.CreatePathMwChain(e, r))
		}
	}
}

func (g *httpProxy) muxRouterHandle(mux *chi.Mux, e *router.Endpoint, r *router.Router) {
	g.mapMethodHandler(mux, r.ReqMethod, r.ReqPath, middleware.CreatePathMwChain(e, r))
}

func (g *httpProxy) mapMethodHandler(mux *chi.Mux, method string, pattern string, handler http.Handler) {
	if strings.Index(pattern, "*") >= 0 {
		length := len(pattern)
		if strings.Index(pattern, "*") == length-1 {
			method = "ALL"
		} else {
			return
		}
	}
	met := strings.ToUpper(method)
	switch met {
	case "ALL":
		mux.Handle(pattern, handler)
	case "GET":
		mux.Method(method, pattern, handler)
	case "POST":
		mux.Method(method, pattern, handler)
	case "PUT":
		mux.Method(method, pattern, handler)
	case "PATCH":
		mux.Method(method, pattern, handler)
	case "DELETE":
		mux.Method(method, pattern, handler)
	case "OPTIONS":
		mux.Method(method, pattern, handler)
	case "HEAD":
		mux.Method(method, pattern, handler)
	default:
		mux.Handle(pattern, handler)
	}
}
