package ratelimiter

import (
	"time"
)

type Rule struct {
	ExpireTime time.Duration
	MaxNum     int
}
