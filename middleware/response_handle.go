package middleware

import (
	"net/http"
	"time"
)

type gatewayResponse struct {
	w            http.ResponseWriter
	status       int
	size         int
	rspBody      []byte
	startTime    time.Time
	pErrorChan   chan error
	pSuccessChan chan bool
	pUseTime     time.Duration
	//pTarget      *core.Target
	reqIp string
}

type gatewayResponseWriter interface {
	Status() int
	Size() int
	RspBody() []byte
	ProxyErrorChan() chan error
	ProxySuccessChan() chan bool
	StartTime() time.Time
	SetProxyUseTime(time.Duration)
	ProxyUseTime() time.Duration
	//SetProxyTarget(*core.Target)
	//ProxyTarget() *core.Target
	SetReqIp(string)
	ReqIp() string
}

func (mw *gatewayResponse) Header() http.Header {
	return mw.w.Header()
}

func (mw *gatewayResponse) Write(b []byte) (int, error) {
	size, err := mw.w.Write(b)
	mw.size += size
	mw.rspBody = append(mw.rspBody, b...)
	return size, err
}

func (mw *gatewayResponse) WriteHeader(s int) {
	mw.w.WriteHeader(s)
	mw.status = s
}

func (mw *gatewayResponse) Status() int {
	return mw.status
}

func (mw *gatewayResponse) Size() int {
	return mw.size
}

func (mw *gatewayResponse) RspBody() []byte {
	return mw.rspBody
}

func (mw *gatewayResponse) ProxyErrorChan() chan error {
	return mw.pErrorChan
}

func (mw *gatewayResponse) ProxySuccessChan() chan bool {
	return mw.pSuccessChan
}

func (mw *gatewayResponse) StartTime() time.Time {
	return mw.startTime
}

func (mw *gatewayResponse) SetProxyUseTime(pUse time.Duration) {
	mw.pUseTime = pUse
}

func (mw *gatewayResponse) ProxyUseTime() time.Duration {
	return mw.pUseTime
}

//func (mw *gatewayResponse) SetProxyTarget(t *def.Target) {
//	mw.pTarget = t
//}

//func (mw *gatewayResponse) ProxyTarget() *core.Target {
//	return mw.pTarget
//}

func (mw *gatewayResponse) SetReqIp(ip string) {
	mw.reqIp = ip
}

func (mw *gatewayResponse) ReqIp() string {
	return mw.reqIp
}
