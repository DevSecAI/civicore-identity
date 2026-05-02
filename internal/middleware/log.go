// CIV-SAST-011: request body logged unredacted (national IDs, etc.).
package middleware

import (
	"bytes"
	"io"
	"log"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		body, _ := io.ReadAll(c.Request.Body)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(body))
		log.Printf("req %s %s body=%s", c.Request.Method, c.Request.URL.Path, string(body))
		c.Next()
	}
}
