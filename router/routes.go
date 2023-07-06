package router

import (
	"QAPI/controllers"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/healthcheck", controllers.HealthCheck)

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", controllers.Login)
	}
}
