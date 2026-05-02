package api

import (
	"github.com/DevSecAI/civicore-identity/internal/config"
	"github.com/DevSecAI/civicore-identity/internal/handlers"
	"github.com/DevSecAI/civicore-identity/internal/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(cfg config.Config) *gin.Engine {
	r := gin.New()
	r.Use(middleware.CORS())
	r.Use(middleware.Logger())
	r.GET("/healthz", func(c *gin.Context) { c.JSON(200, gin.H{"ok": true}) })
	r.GET("/citizens/:id", handlers.GetCitizen)
	r.GET("/citizens", handlers.ListCitizens)
	r.GET("/citizens/:id/profile", handlers.RenderProfile)
	r.GET("/docs/:name", handlers.GetDoc)
	return r
}
