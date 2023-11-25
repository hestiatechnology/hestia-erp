package routes

import (
	"hestia/api/middleware"
	"hestia/api/models"
	"hestia/api/utils"
	"log"

	"github.com/gin-gonic/gin"
)

func ClientGet(ctx *gin.Context) {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var clients []models.Client
	rows, err := db.Query(
		ctx.Request.Context(),
		"SELECT id, name, code, vat_id, street, postal_code, locality, country FROM sales.client WHERE company_id = $1",
		ctx.GetHeader("X-Company-Id"),
	)

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var c models.Client
		err := rows.Scan(&c.Id, &c.Name, &c.Code, &c.VatId, &c.Street, &c.PostalCode, &c.Locality, &c.Country)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, c)
	}

	rows.Close()

	ctx.JSON(200, clients)
}

func ClientRoutes(r *gin.Engine) {
	client := r.Group("/client", middleware.BearerAuthenticate(), middleware.CompanyId())

	// /company
	client.GET("", ClientGet)

}
