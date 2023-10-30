package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func MethodNotAllowed(ctx *gin.Context) {
	ctx.AbortWithStatus(http.StatusMethodNotAllowed)
}
