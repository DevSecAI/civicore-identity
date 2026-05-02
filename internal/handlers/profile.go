// CIV-SAST-004: writes citizen-controlled name as raw HTML.
package handlers

import "github.com/gin-gonic/gin"

func RenderProfile(c *gin.Context) {
	id := c.Param("id")
	name := c.Query("name")
	c.Data(200, "text/html",
		[]byte("<html><body><h1>Profile of " + name + " (" + id + ")</h1></body></html>"))
}
