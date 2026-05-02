// CIV-SAST-006: SSRF — http.Get on caller URL.
package services

import (
	"io"
	"net/http"
)

func FetchExternal(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil { return "", err }
	defer resp.Body.Close()
	b, _ := io.ReadAll(resp.Body)
	return string(b), nil
}
