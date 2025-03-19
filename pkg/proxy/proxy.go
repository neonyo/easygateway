package proxy

import "sync"

type Proxy struct {
	sync.RWMutex
	//ngin   *engine
}

func NewProxy() *Proxy {
	return &Proxy{}
}
