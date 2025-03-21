package middleware

import (
	"net/http"
	"net/url"
	"time"
)

type Gateway struct {
	Baser
}

func (p *Gateway) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if g, ok := w.(gatewayResponseWriter); ok {
		endpoint := p.GetEndpoint()
		route := p.GetRouter()
		remote, _ := url.Parse(endpoint.Addr)
		proxy := NewSingleHostReverseProxy(remote, route)
		proxy.ErrorHandler = func(rw http.ResponseWriter, request *http.Request, e error) {
			g.ProxyErrorChan() <- e
		}
		proxy.SuccessHandler = func(rw http.ResponseWriter, request *http.Request) {
			g.ProxySuccessChan() <- true
		}
		proxy.ServeHTTP(w, r)
		g.SetProxyUseTime(time.Now().Sub(time.Now()))
	}
}
