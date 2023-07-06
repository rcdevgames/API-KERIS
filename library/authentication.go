package library

import (
	"QAPI/configs"
	"QAPI/entities"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexsasharegan/dotenv"
	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
)

func Authorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		errenv := dotenv.Load()
		if errenv != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Kesalahan Server",
			})
			ctx.Abort()
			return
		}

		contentType := ctx.Request.Header.Get("Content-Type")
		bearerToken := ctx.Request.Header.Get("Authorization")
		if len(contentType) > 0 {
			var allowedToken = false
			if len(strings.Split(bearerToken, "")) >= 2 {
				cleanToken := strings.Split(bearerToken, " ")[1]

				token, err := jwt.Parse(cleanToken, func(t *jwt.Token) (interface{}, error) {
					if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
						return nil, fmt.Errorf("unexpected signing method")
					}
					return []byte(os.Getenv("APP_KEY")), nil
				})

				if err == nil {
					claims, ok := token.Claims.(jwt.MapClaims)
					if ok && token.Valid {
						expired := getTokenRemainingValidity(claims["exp"])
						if expired < 0 {
							allowedToken = false
						} else {
							var id = claims["id"]
							var user *entities.Users
							data := configs.DB.Find(&user, "id = ?", id)
							if data.RowsAffected > 0 {
								ctx.Set("id_user", strconv.Itoa(user.Id))
								ctx.Set("id_device", user.DeviceId)
								allowedToken = true
							}
						}
					}
				}

			}

			if !allowedToken {
				ctx.JSON(http.StatusUnauthorized, gin.H{
					"success": false,
					"message": "Token tidak valid",
				})
				ctx.Abort()
				return
			}
		} else {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Header wajib JSON!",
			})
			ctx.Abort()
			return
		}
	}
}

func getTokenRemainingValidity(timestamp interface{}) int {
	if validity, ok := timestamp.(float64); ok {
		tm := time.Unix(int64(validity), 0)
		remainer := tm.Sub(time.Now())
		if remainer > 0 {
			return int(remainer.Seconds())
		}
	}
	return -1
}
