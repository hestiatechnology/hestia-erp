package routes

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func CompanyGet(ctx *gin.Context) {
	connStr := "postgres://postgres:alexis27@localhost/erp?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id FROM companies.company")
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
	company.GET("", CompanyGet)

}
