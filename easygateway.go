package easygateway

type EasyGateway struct {
	Endpoint     map[string]Endpoint
	whitEndpoint func()
}

func (g *EasyGateway) SetEndpoint(f func()) {
	g.whitEndpoint = f
}

func (g *EasyGateway) Start() {

}
