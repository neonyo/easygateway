package util

import (
	"bytes"
	"compress/gzip"
	"io"
	"net/http"
	"strings"
)

func IsGzip(header http.Header) bool {
	needGzip := false
	if headerVal, ok := header["Content-Encoding"]; ok {
		for _, data := range headerVal {
			if strings.Index(data, "gzip") != -1 {
				needGzip = true
			}
		}
	}
	return needGzip
}

func GzipByteToString(p []byte) string {
	if len(p) == 0 {
		return ""
	}
	buf := bytes.NewBuffer([]byte{})
	buf.Write(p)
	read, _ := gzip.NewReader(buf)
	data, _ := io.ReadAll(read)
	return string(data)
}
