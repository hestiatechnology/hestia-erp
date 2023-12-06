package routes

import (
	"hestia/api/logger"
	"hestia/api/middleware"
	"hestia/api/models"
	"hestia/api/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

// /clients GET
func clientsGet(ctx *gin.Context) {
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

	db, err := utils.GetDbPoolConn()
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

	var count int
	for rows.Next() {
		var c models.Client
		count++
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

	if count == 0 {
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.JSON(200, clients)
}

// /clients POST
func clientsPost(ctx *gin.Context) {

	var newClient models.NewClient
	if err := ctx.ShouldBindJSON(&newClient); err != nil {

		logger.Error.Println("Error while binding JSON: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Input validation failed, check documentation for correct input types",
		})
		return
	}

	db, err := utils.GetDbPoolConn()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(ctx.Request.Context())
	if err != nil {
		logger.Error.Println("Error while starting transaction: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while creating client",
		})
		return
	}

	defer tx.Rollback(ctx.Request.Context())

	clientCodeExists := tx.QueryRow(
		ctx.Request.Context(),
		"SELECT COUNT(id) FROM sales.client WHERE company_id = $1 AND code = $2",
		ctx.GetHeader("X-Company-Id"),
		newClient.Code,
	)

	var clientCodeCount int
	err = clientCodeExists.Scan(&clientCodeCount)
	if err != nil {
		logger.Error.Println("Error while checking if client code exists: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while creating client",
		})
		return
	}

	if clientCodeCount > 0 {
		logger.Warning.Println("Client code already exists")
		ctx.AbortWithStatusJSON(http.StatusConflict, models.ErrorMessage{
			Message: "A client with that code already exists",
		})
		return
	}

	clientId := tx.QueryRow(
		ctx.Request.Context(),
		"INSERT INTO sales.client (company_id, name, code, vat_id, street, postal_code, locality, country) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		ctx.GetHeader("X-Company-Id"),
		newClient.Name,
		newClient.Code,
		newClient.VatId,
		newClient.Street,
		newClient.PostalCode,
		newClient.Locality,
		newClient.Country,
	)

	var id uuid.UUID
	err = clientId.Scan(&id)
	if err != nil {
		logger.Error.Println("Error while inserting client: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while creating client",
		})
		return
	}

	err = tx.Commit(ctx.Request.Context())
	if err != nil {
		logger.Error.Println("Error while committing transaction: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while creating client",
		})
		return
	}

	ctx.Header("Location", "/clients/"+id.String())
	ctx.JSON(201, gin.H{
		"id": id,
	})
}

// /clients/:id GET
func clientsIdGet(ctx *gin.Context) {
	id := ctx.Param("id")

	db, err := utils.GetDbPoolConn()
	if err != nil {
		log.Fatal(err)
	}

	var client models.Client
	row := db.QueryRow(
		ctx.Request.Context(),
		"SELECT id, name, code, vat_id, street, postal_code, locality, country FROM sales.client WHERE id = $1 AND company_id = $2",
		id,
		ctx.GetHeader("X-Company-Id"),
	)

	err = row.Scan(&client.Id, &client.Name, &client.Code, &client.VatId, &client.Street, &client.PostalCode, &client.Locality, &client.Country)
	// If no rows are returned, err will be set to sql.ErrNoRows
	if err == pgx.ErrNoRows {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if err != nil {
		logger.Error.Println("Error while scanning client: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while getting client",
		})
		return
	}

	ctx.JSON(200, client)
}

// /clients/:id PUT
/*
func clientsIdPut(ctx *gin.Context) {
	id := ctx.Param("id")

	var newClient models.NewClient
	if err := ctx.ShouldBindJSON(&newClient); err != nil {

		logger.Error.Println("Error while binding JSON: ", err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, models.ErrorMessage{
			Message: "Input validation failed, check documentation for correct input types",
		})
		return
	}

	db, err := utils.GetDbPoolConn()
	if err != nil {
		log.Fatal(err)
	}

	tx, err := db.Begin(ctx.Request.Context())
	if err != nil {
		logger.Error.Println("Error while starting transaction: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while updating client",
		})
		return
	}

	defer tx.Rollback(ctx.Request.Context())

	clientCodeExists := tx.QueryRow(
		ctx.Request.Context(),
		"SELECT COUNT(id) FROM sales.client WHERE company_id = $1 AND code = $2 AND id != $3",
		ctx.GetHeader("X-Company-Id"),
		newClient.Code,
		id,
	)

	var clientCodeCount int
	err = clientCodeExists.Scan(&clientCodeCount)
	if err != nil {
		logger.Error.Println("Error while checking if client code exists: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while updating client",
		})
		return
	}

	if clientCodeCount > 0 {
		logger.Warning.Println("Client code already exists")
		ctx.AbortWithStatusJSON(http.StatusConflict, models.ErrorMessage{
			Message: "A client with that code already exists",
		})
		return
	}

	clientId := tx.QueryRow(
		ctx.Request.Context(),
		"UPDATE sales.client SET name = $1, code = $2, vat_id = $3, street = $4, postal_code = $5, locality = $6, country = $7 WHERE id = $8 AND company_id = $9 RETURNING id",
		newClient.Name,
		newClient.Code,
		newClient.VatId,
		newClient.Street,
		newClient.PostalCode,
		newClient.Locality,
		newClient.Country,
		id,
		ctx.GetHeader("X-Company-Id"),
	)

	var uuid uuid.UUID
	err = clientId.Scan(&uuid)
	if err != nil {
		logger.Error.Println("Error while updating client: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while updating client",
		})
		return
	}

	err = tx.Commit(ctx.Request.Context())
	if err != nil {
		logger.Error.Println("Error while committing transaction: ", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, models.ErrorMessage{
			Message: "Error while updating client",
		})
		return
	}

	ctx.Header("Location", "/clients/"+uuid.String())
	ctx.JSON(200, gin.H{
		"id": uuid,
	})
}
*/
func ClientsRoutes(r *gin.Engine) {
	client := r.Group("/clients", middleware.BearerAuthenticate(), middleware.CompanyId())

	// /client
	client.GET("", clientsGet)
	client.POST("", clientsPost)

	// /client/:id
	clientId := client.Group("/:id")
	clientId.GET("", clientsIdGet)
	//clientId.PUT("", clientsIdPut)

}
