package middleware

import (
	"github.com/neonyo/easygateway/pkg/util"
	"net"
	"net/http"
)

type BlackIpsMw struct {
	Baser
}

func (mw *BlackIpsMw) Init() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(mw.GetEndpoint().BlackIps) > 0 {
				ip := realIP(r)
				hgwResponse := w.(*gatewayResponse)
				hgwResponse.SetReqIp(ip)
				if ok := util.Contains(mw.GetEndpoint().BlackIps, ip); ok {
					w.WriteHeader(http.StatusForbidden)
					return
				}
			}
			next.ServeHTTP(w, r)
		})
	}
}

const (
	XForwardedFor = "X-Forwarded-For"
	XRealIP       = "X-Real-IP"
)

func realIP(req *http.Request) string {
	remoteAddr := req.RemoteAddr
	if ip := req.Header.Get(XRealIP); ip != "" {
		remoteAddr = ip
	} else if ip = req.Header.Get(XForwardedFor); ip != "" {
		remoteAddr = ip
	} else {
		remoteAddr, _, _ = net.SplitHostPort(remoteAddr)
	}

	if remoteAddr == "::1" {
		remoteAddr = "127.0.0.1"
	}

	return remoteAddr
}
