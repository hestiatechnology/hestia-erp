package main

import (
	"hestia/api/middleware"
	"hestia/api/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.New()
	r.RedirectTrailingSlash = false
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.RequestIdMiddleware())
	r.Use(middleware.DontCache())
	routes.SetRoutes(r)
	return r
}

func main() {
	r := SetupRouter()
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
