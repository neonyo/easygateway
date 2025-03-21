package middleware

import (
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/neonyo/easygateway/pkg/errorx"
	"net/http"
)

type BreakerMw struct {
	Baser
}

func (m *BreakerMw) Init() func(http.Handler) http.Handler {
	endpoint := m.GetEndpoint()
	route := m.GetRouter()
	cmdConf := hystrix.CommandConfig{
		MaxConcurrentRequests:  route.CircuitBreakerRequest,
		RequestVolumeThreshold: route.CircuitVolumeThreshold,
		ErrorPercentThreshold:  route.CircuitBreakerPercent,
		SleepWindow:            route.CircuitSleepWindow,
	}
	if route.CircuitBreakerTimeout > 0 {
		cmdConf.Timeout = route.CircuitBreakerTimeout
	}
	hystrix.ConfigureCommand(fmt.Sprint(route.ReqPath), cmdConf)
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			gw := w.(gatewayResponseWriter)
			_ = hystrix.Do(fmt.Sprint(r.URL.Path), func() (err error) {
				next.ServeHTTP(w, r)
				select {
				case <-gw.ProxySuccessChan():
					return nil
				case err = <-gw.ProxyErrorChan():
					return err
				}
			}, func(err error) error {
				_, _ = w.Write(errorx.New(err, endpoint))
				return nil
			})
		})
	}
}
