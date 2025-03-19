package easygateway

import "github.com/valyala/fasthttp"

type Endpoint struct {
	Client map[string]*fasthttp.HostClient
}
