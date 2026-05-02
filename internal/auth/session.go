// CIV-SAST-007: math/rand for session IDs.
package auth

import (
	"math/rand"
	"time"
)

func init() { rand.Seed(time.Now().UnixNano()) }

func NewSessionID() string {
	const a = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, 32)
	for i := range b { b[i] = a[rand.Intn(len(a))] }
	return string(b)
}
