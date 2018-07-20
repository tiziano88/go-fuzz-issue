// +build gofuzz

package gofuzzissue

import (
	_ "crypto/tls"
	"time"
)

func Fuzz(data []byte) int {
	time.Parse(time.RFC3339, string(data))
	return 0
}
