package router

import (
	"QAPI/controllers"
	"QAPI/library"

	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/healthcheck", controllers.HealthCheck)

	authRoute := router.Group("/auth")
	{
		authRoute.POST("/login", controllers.Login)
	}

	authorizedRoute := router.Group("/v1")
	authorizedRoute.Use(library.Authorize())
	{
		merchantRoute := authorizedRoute.Group("/merchant")
		{
			merchantRoute.GET("/", controllers.GetDetail)
			merchantRoute.POST("/register", controllers.Register)
		}

		trxRoute := authorizedRoute.Group("/trx")
		{
			trxRoute.POST("/", controllers.GenerateQRIS)
		}
	}
}
