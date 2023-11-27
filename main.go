package main

import (
	"hestia/api/logger"
	"hestia/api/middleware"
	"hestia/api/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	if os.Getenv("HESTIA_ENV") != "production" {
		r.Use(gin.Logger())
	}
	if os.Getenv("HESTIA_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}
	r.UseH2C = true
	r.ForwardedByClientIP = true
	r.RedirectTrailingSlash = false
	r.Use(gin.Recovery())
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.DontCache())
	r.HandleMethodNotAllowed = true
	r.NoMethod(middleware.MethodNotAllowed())
	r.NoRoute(middleware.NotFound())
	routes.SetRoutes(r)
	return r
}

func main() {
	logger.InitLogger()
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
