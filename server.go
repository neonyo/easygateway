package gw

import (
	"github.com/neonyo/gw/router"
	"net/http"
)

type Server struct {
	proxy       *proxy
	options     []Option
	httpOn      bool
	grpcOn      bool
	websocketOn bool
}

func New() *Server {
	return &Server{
		proxy: newProxy(),
	}
}

// RegisterHttp 设置http设置
func (g *Server) RegisterHttp(conf HttpConfig, f func(r *http.Request) string) {
	g.httpOn = true
	g.proxy.cfg.httpConfigOption = append(g.proxy.cfg.httpConfigOption, withHttpConf(conf), withEndpointSwitch(f))
}

// AddEndpoints 新增服务
func (g *Server) AddEndpoints(rs []*router.Endpoint) {
	g.proxy.httpProxy.addEndpoints(rs)
}

// UpdateEndpoint 更新服务
func (g *Server) UpdateEndpoint(r *router.Endpoint) {
	g.proxy.httpProxy.updateEndpoint(r)
}

// DelEndpoint 删除服务
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

func (g *Server) Start(addr string) error {
	if g.httpOn {
		g.proxy.withHttp()
	}
	return g.proxy.start(addr)
}
