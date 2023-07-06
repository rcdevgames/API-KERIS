package controllers

import (
	"QAPI/entities"
	"QAPI/library"
	"QAPI/logger"
	"QAPI/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	var register_model entities.MerchantInsert
	if errbody := ctx.BindJSON(&register_model); errbody != nil {
		logger.Log.Err(errbody).Msg("Error Body Register Merchant:")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input Body Invalid",
		})
		return
	}

	// Check Login Dulu
	qris := library.InitQris(register_model.QrisEmail, register_model.QrisPassword)
	data_merchant, err := qris.Merchant()
	if err != nil {
		logger.Log.Err(err).Msg("Gagal Req Merchant")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	data, dataerr := library.Encode(register_model.QrisData)
	if dataerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": dataerr.Error(),
		})
		return
	}
	email, emailerr := library.Encode(register_model.QrisEmail)
	if emailerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": emailerr.Error(),
		})
		return
	}
	pass, passerr := library.Encode(register_model.QrisPassword)
	if passerr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": passerr.Error(),
		})
		return
	}
	id, err := strconv.Atoi(ctx.GetString("id_user"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	insert_data := entities.MerchantInsert{
		UserID:       id,
		QrisData:     data,
		QrisEmail:    email,
		QrisPassword: pass,
		QrisName:     data_merchant.Name,
		QrisMnid:     data_merchant.MNID,
	}
	if res := models.InsertMerchant(insert_data, ctx.GetString("id_device")); !res {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal Registrasi Merchant, mohon kontak admin.",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Registrasi Merchant Berhasil",
	})
}

func GetDetail(ctx *gin.Context) {

}
