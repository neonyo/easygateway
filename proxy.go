package gw

import (
	"sync"
)

type (
	proxy struct {
		sync.RWMutex
		cfg       *config
		httpProxy *httpProxy
	}
)

func newProxy() *proxy {
	return &proxy{
		cfg:       new(config),
		httpProxy: &httpProxy{},
	}
}

func (p *proxy) withHttp() {
	p.httpProxy.conf = NewHttpConfig(p.cfg.httpConfigOption...)
}
