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
	var limitOffset models.LimitOffset
	if err := ctx.ShouldBindJSON(&limitOffset); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Missing limit and offset",
		})
		return
	}

	if limitOffset.Limit <= 0 || limitOffset.Offset < 0 || limitOffset.Limit > 100 || limitOffset.Offset > 100 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Invalid limit and offset",
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
		limitOffset.Limit,
		limitOffset.Offset,
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
