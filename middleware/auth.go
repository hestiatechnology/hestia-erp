package middleware

import (
	"hestia/api/models"
	"hestia/api/utils"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BearerAuthenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorMessage{
				Message: "Missing header Authorization",
			})
			return
		}

		// Extract the token from the Authorization header
		// The header format is "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Connect to DB and check if the token exists
		db, err := utils.GetDbPoolConn()
		if err != nil {
			log.Fatal(err)
		}

		row := db.QueryRow(ctx.Request.Context(), "SELECT count(id) FROM users.users_session WHERE id = $1", token)

		var count int
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// If the token is valid, continue with the request
		ctx.Next()
	}

}

func CompanyId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		companyId := ctx.GetHeader("X-Company-Id")
		if companyId == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorMessage{
				Message: "Missing header X-Company-Id",
			})
			return
		}

		// Validate the token
		// Connect to DB and check if the company exists and the user has access to it via the table user_company
		db, err := utils.GetDbPoolConn()
		if err != nil {
			log.Fatal(err)
		}

		userId, err := utils.GetUserId(ctx.Request.Context(), utils.GetSessionId(ctx.GetHeader("Authorization")))
		if err != nil {
			log.Fatal(err)
		}

		row := db.QueryRow(
			ctx.Request.Context(),
			"SELECT count(company_id) FROM users.user_company WHERE user_id = $1 AND company_id = $2",
			userId,
			companyId,
		)
		var count int
		err = row.Scan(&count)
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}
		// If the token is valid, continue with the request
		ctx.Next()
	}

}
