package middleware

import (
	"hestia/api/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func RequestIdMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		reqId := uuid.New().String()
		ctx.Writer.Header().Set("X-Request-Id", reqId)
		ctx.Next()
	}
}

func DontCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Cache-Control", "no-store")
		ctx.Next()
	}
}

func MethodNotAllowed() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusMethodNotAllowed)
	}
}

func NotFound() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatus(http.StatusNotFound)
	}
}

func BearerAuthenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorMessage{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Missing header Authorization",
			})
			return
		}

		// Extract the token from the Authorization header
		// The header format is "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		// In this example, we're just checking if it matches a predefined token
		if token != "expected_token" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, models.ErrorMessage{
				Code:    http.StatusUnauthorized,
				Status:  "Unauthorized",
				Message: "Invalid token",
			})
			return
		}

		// If the token is valid, continue with the request
		ctx.Next()
	}

}
