package main

import (
	"hestia/api/middleware"
	"hestia/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.Use(middleware.RequestIdMiddleware())
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	routes.SetRoutes(r)
	//routes.SetRoutes(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run(":8080")
}
