package gw

import "net/http"

type HttpConfig struct {
	Addr       string
	Telemetry  bool
	hostSwitch func(r *http.Request) string
}

func NewHttpConfig(options ...HttpOption) HttpConfig {
	var c HttpConfig
	for _, option := range options {
		c = option.applyHttpStart(c)
	}
	return c
}

type HttpOption interface {
	applyHttpStart(HttpConfig) HttpConfig
}
type httpOptionFunc func(HttpConfig) HttpConfig

func (fn httpOptionFunc) applyHttpStart(c HttpConfig) HttpConfig {
	return fn(c)
}

func withHttpConf(httpConf HttpConfig) HttpOption {
	return httpOptionFunc(func(c HttpConfig) HttpConfig {
		return httpConf
	})
}

// WithEndpointSwitch 设置服务地址选择
func withEndpointSwitch(fn func(r *http.Request) string) HttpOption {
	return httpOptionFunc(func(c HttpConfig) HttpConfig {
		c.hostSwitch = fn
		return c
	})
}
