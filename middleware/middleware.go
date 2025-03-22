package middleware

import (
	"github.com/justinas/alice"
	"github.com/neonyo/gw/router"
	"net/http"
)

type HgwMiddleWare interface {
	Init() func(http.Handler) http.Handler
}

func mwList(chain *[]alice.Constructor, hd func(http.Handler) http.Handler) bool {
	*chain = append(*chain, hd)
	return true
}

func CreatePathMwChain(endpoint *router.Endpoint, router *router.Router) http.Handler {
	baseMw := &Base{Endpoint: endpoint, Router: router}
	//baseMw.SetMt(core.NewDomainPathMetrics(domain, path))
	return createMwChain(baseMw)
}

func createMwChain(base *Base) http.Handler {
	var chainArray []alice.Constructor
	mwList(&chainArray, base.Init())
	mwList(&chainArray, (&RecoverMw{base}).Init())
	mwList(&chainArray, (&BlackIpsMw{base}).Init())
	//mwList(&chainArray, (&MetricsMw{base}).Init())
	//mwList(&chainArray, (&RequestCopyMw{base}).Init())
	//mwList(&chainArray, (&LoggerMw{base}).Init())
	//mwList(&chainArray, (&RateLimiterMw{base}).Init())
	mwList(&chainArray, (&BreakerMw{base}).Init())
	gw := &Gateway{base}
	hl := alice.New(chainArray...).Then(gw)
	return hl
}
