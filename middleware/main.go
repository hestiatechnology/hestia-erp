package middleware

import (
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

func BearerAuthenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authHeader := ctx.GetHeader("Authorization")
		if authHeader == "" {
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Extract the token from the Authorization header
		// The header format is "Bearer <token>"
		token := strings.TrimPrefix(authHeader, "Bearer ")

		// Validate the token
		// In this example, we're just checking if it matches a predefined token
		if token != "expected_token" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// If the token is valid, continue with the request
		ctx.Next()
	}

}
