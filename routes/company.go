package routes

import (
	"hestia/api/utils"

	"github.com/gin-gonic/gin"
)

func CompanyGet(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "companies",
	})
}

func SetRoutes(r *gin.Engine) {
	company := r.Group("/company")

	// /company
	company.GET("/", CompanyGet)
	company.POST("/", utils.MethodNotAllowed)
	company.PUT("/", utils.MethodNotAllowed)
	company.DELETE("/", utils.MethodNotAllowed)

}
