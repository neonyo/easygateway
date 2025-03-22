package errorx

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/neonyo/gw/router"
	"net"
	"net/http"
)

type CustomError struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func New(err error, endpoint *router.Endpoint) []byte {
	if err == nil {
		return []byte("")
	}
	var msg string
	var opError *net.OpError
	var hystrixErr hystrix.CircuitError
	switch {
	case errors.As(err, &opError):
		msg = fmt.Sprintf("%s服务错误：%s %s", endpoint.Name, opError.Op, opError.Err.Error())
	case errors.As(err, &hystrixErr):
		switch hystrixErr.Message {
		case "timeout":
			msg = "请求超时"
		case "circuit open":
			msg = "请求已断开"
		case "max concurrency":
			msg = "请求次数过多，请歇会再操作"
		default:
			msg = hystrixErr.Message
		}
	default:
		msg = err.Error()
	}
	b, _ := json.Marshal(CustomError{
		Code: http.StatusBadRequest,
		Msg:  msg,
	})
	return b
}
