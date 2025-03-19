package rest

import "github.com/valyala/fasthttp"

type (
	Server struct {
		endpoint *Endpoint
	}
)

func MustNewServer() *Server {
	return &Server{
		endpoint: newEndpoint(),
	}
}

func (s *Server) AddEndpoint(fastClient *fasthttp.HostClient) {
	s.endpoint.addEndpoint(fastClient)
}

func (s *Server) Start() {
	s.endpoint.start()
}
