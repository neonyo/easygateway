package easygateway

import (
	"github.com/neonyo/easygateway/router"
)

type Server struct {
	proxy *proxy
}

func New(c Conf) *Server {
	return &Server{
		proxy: newProxy(proxyConfig{
			Addr:      c.Addr,
			Telemetry: c.Telemetry,
		}),
	}
}

// AddEndpoints 新增代理服务
func (g *Server) AddEndpoints(rs []*router.Endpoint) {
	g.proxy.httpProxy.addEndpoints(rs)
}

// UpdateEndpoint 更新单个代理服务
func (g *Server) UpdateEndpoint(r *router.Endpoint) {
	g.proxy.httpProxy.updateEndpoint(r)
}

// DelEndpoint 删除单个代理服务
func (g *Server) DelEndpoint(r *router.Endpoint) {
	g.proxy.httpProxy.delEndpoint(r)
}

// AddRouters 添加路由
func (g *Server) AddRouters(e *router.Endpoint, rs []*router.Router) {
	g.proxy.httpProxy.addRouter(e, rs)
}

// UpdateRouter 更新路由
func (g *Server) UpdateRouter(e *router.Endpoint) {
	g.proxy.httpProxy.updateRouter(e)
}

// WithEndpointOption 设置判断服务站点
func (g *Server) WithEndpointOption(endpointFn endpointOption) {
	g.proxy.withEndpointOption(endpointFn)
}

func (g *Server) Start() error {
	return g.proxy.start()
}
