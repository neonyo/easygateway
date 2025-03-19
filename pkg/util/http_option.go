package util

import "time"

// HTTPOption http client option
type HTTPOption struct {
	// Maximum number of connections which may be established to server
	MaxConn int
	// MaxConnDuration Keep-alive connections are closed after this duration.
	MaxConnDuration time.Duration
	// MaxIdleConnDuration Idle keep-alive connections are closed after this duration.
	MaxIdleConnDuration time.Duration
	// ReadBufferSize Per-connection buffer size for responses' reading.
	ReadBufferSize int
	// WriteBufferSize Per-connection buffer size for requests' writing.
	WriteBufferSize int
	// ReadTimeout Maximum duration for full response reading (including body).
	ReadTimeout time.Duration
	// WriteTimeout Maximum duration for full request writing (including body).
	WriteTimeout time.Duration
	// MaxResponseBodySize Maximum response body size.
	MaxResponseBodySize int
	// DisableHeaderNamesNormalizing disable normalizing the header name
	DisableHeaderNamesNormalizing bool
}
