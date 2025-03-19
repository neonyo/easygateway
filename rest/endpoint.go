package rest

import "github.com/valyala/fasthttp"

type Endpoint struct {
	fastClient []*fasthttp.HostClient
}

func newEndpoint() *Endpoint {
	return &Endpoint{}
}

func (e *Endpoint) addEndpoint(fastClient *fasthttp.HostClient) {
	e.fastClient = append(e.fastClient, fastClient)
}

func (e *Endpoint) start() {

}
