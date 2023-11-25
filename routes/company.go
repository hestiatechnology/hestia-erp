package routes

import (
	"hestia/api/middleware"
	"hestia/api/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func CompanyGet(ctx *gin.Context) {

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	db.Ping(ctx.Request.Context())
	rows, err := "db.QueryRow()", nil
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"rows": rows,
	})
}

func CompanyRoutes(r *gin.Engine) {
	company := r.Group("/company", middleware.CompanyId())

	// /company
	company.GET("", CompanyGet)

}
