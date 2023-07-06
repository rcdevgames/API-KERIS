package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HealthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Always ON Guys :)",
	})
}
