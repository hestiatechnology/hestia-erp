package routes

import (
	"database/sql"
	"hestia/api/middleware"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CompanyGet(ctx *gin.Context) {

	connStr := "postgres://postgres:alexis27@localhost/erp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := "db.QueryRow()", nil
	if err != nil {
		log.Fatal(err)
	}
	ctx.JSON(200, gin.H{
		"rows": rows,
	})
}

func CompanyRoutes(r *gin.Engine) {
	company := r.Group("/company")

	// /company
	company.GET("", middleware.CompanyId(), CompanyGet)

}
