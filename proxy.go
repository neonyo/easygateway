package easygateway

import (
	"net/http"
	"sync"
)

type proxyConfig struct {
	Addr      string
	Telemetry bool
}

type (
	endpointOption func(r *http.Request) string
	proxy          struct {
		sync.RWMutex
		cfg proxyConfig
		//endpoints []*router.Endpoint
		httpProxy *httpProxy
	}
)

func newProxy(cfg proxyConfig) *proxy {
	return &proxy{
		cfg: cfg,
		httpProxy: &httpProxy{
			Telemetry: cfg.Telemetry,
		},
	}
}

func (p *proxy) withEndpointOption(fn endpointOption) {
	p.httpProxy.endpointOption = fn
}
