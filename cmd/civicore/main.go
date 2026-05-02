package main

import (
	"github.com/DevSecAI/civicore-identity/internal/api"
	"github.com/DevSecAI/civicore-identity/internal/config"
)

func main() {
	cfg := config.Load()
	r := api.NewRouter(cfg)
	r.Run(cfg.Listen)
}
