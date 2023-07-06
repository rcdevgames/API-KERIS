package controllers

import (
	"QAPI/entities"
	"QAPI/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context) {
	var login_model entities.LoginModel
	if errbody := ctx.BindJSON(&login_model); errbody != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input Body Invalid",
		})
		return
	}

user_check:
	user := models.GetUserByIdDevice(login_model.DeviceId)
	if user == nil {
		user_insert := entities.UserInsert{
			DeviceId:    login_model.DeviceId,
			OnesignalId: login_model.OnesignalId,
			CreatedBy:   login_model.DeviceId,
		}
		reg_user := models.CreateUser(user_insert)
		if !reg_user {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Cannot Register User",
			})
			return
		}

		goto user_check
		return
	} else {

	}
}
