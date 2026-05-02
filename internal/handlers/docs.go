// CIV-SAST-005: path traversal via filepath.Join with user input.
package handlers

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

const docsRoot = "/var/civicore/docs"

func GetDoc(c *gin.Context) {
	name := c.Param("name")
	path := filepath.Join(docsRoot, name)
	c.File(path)
}
