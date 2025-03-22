package gw

type Conf struct {
	ServerName string
	Addr       string
	Telemetry  bool
	Mid        string
}

type config struct {
	ServerName       string
	httpConfigOption []HttpOption
}

type Option interface {
	apply(*config)
}
type optionFunc func(*config)

func (o optionFunc) apply(c *config) {
	o(c)
}

func newConfig(opts ...Option) *config {
	c := &config{}
	for _, opt := range opts {
		opt.apply(c)
	}
	return c
}

//func withHttpOption(ops ...HttpOption) Option {
//	return optionFunc(func(c *config) {
//		c.httpConfigOption = append(c.httpConfigOption, ops...)
//	})
//}

func WithWebSocketOption() {

}

func WithGrpcOption() {

}
