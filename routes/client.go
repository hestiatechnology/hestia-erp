package routes

import (
	"hestia/api/logger"
	"hestia/api/middleware"
	"hestia/api/models"
	"hestia/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ClientGet(ctx *gin.Context) {
	type FilterOptions struct {
		models.Client
		models.LimitOffset
	}
	var filterOptions FilterOptions
	if err := ctx.ShouldBindJSON(&filterOptions); err != nil {

		logger.Error.Println("Error while binding JSON: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Input validation failed, check documentation for correct input types",
		})
		return
	}

	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

	var clients []models.Client
	rows, err := db.Query(
		ctx.Request.Context(),
		"SELECT id, name, code, vat_id, street, postal_code, locality, country FROM sales.client WHERE company_id = $1 LIMIT $2 OFFSET $3",
		ctx.GetHeader("X-Company-Id"),
		filterOptions.Limit,
		filterOptions.Offset,
	)

	if err != nil {
		logger.Error.Println("Error while querying clients: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while fetching clients",
		})
		return
	}

	defer rows.Close()

	for rows.Next() {
		var c models.Client
		err := rows.Scan(&c.Id, &c.Name, &c.Code, &c.VatId, &c.Street, &c.PostalCode, &c.Locality, &c.Country)
		if err != nil {
			logger.Error.Println("Unable to scan rows into Client model, error: ", err)
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
				Message: "Error while getting clients",
			})
			return
		}
		clients = append(clients, c)
	}

	ctx.JSON(200, clients)
}

func ClientRoutes(r *gin.Engine) {
	client := r.Group("/client", middleware.BearerAuthenticate(), middleware.CompanyId())

	// /client
	client.GET("", ClientGet)

}
