package controllers

import (
	"QAPI/entities"
	"QAPI/logger"
	"QAPI/models"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
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
	} else {
		if user.IsActive != 1 {
			ctx.JSON(http.StatusForbidden, gin.H{
				"success": false,
				"message": "Akun anda sudah dinonaktifkan.",
			})
			return
		}
		claims := jwt.MapClaims{
			"uid":       strconv.Itoa(user.Id),
			"exp":       time.Now().Add(time.Minute * 30).Unix(),
			"iss":       "dev.rcdevgames.net",
			"id":        strconv.Itoa(user.Id),
			"device_id": user.DeviceId,
		}
		sign := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		token, errjwt := sign.SignedString([]byte(os.Getenv("APP_KEY")))
		if errjwt != nil {
			logger.Log.Err(errjwt).Msg("Error JWT " + user.DeviceId + ":")
			ctx.JSON(http.StatusBadRequest, gin.H{
				"success": false,
				"message": "Terjadi kesalahan saat login, mohon coba kebali",
			})
			return
		}
		models.UpdateLastLogin(user.Id)
		ctx.JSON(http.StatusOK, gin.H{
			"success": true,
			"message": "Login Berhasil",
			"data": entities.LoginResponse{
				AccessToken: token,
			},
		})
	}
}
