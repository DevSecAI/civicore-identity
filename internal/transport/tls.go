// CIV-SAST-008: TLS InsecureSkipVerify=true.
package transport

import (
	"crypto/tls"
	"net/http"
)

func Client() *http.Client {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	return &http.Client{Transport: tr}
}
