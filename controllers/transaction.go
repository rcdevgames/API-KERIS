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

func GenerateQRIS(ctx *gin.Context) {
	var trx_model entities.TransactionInput
	if errbody := ctx.BindJSON(&trx_model); errbody != nil {
		logger.Log.Err(errbody).Msg("Error Body Generate QRIS:")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Input Body Invalid",
		})
		return
	}

	if trx_model.Amount < 1 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Nomimal minimal Rp. 1",
		})
		return
	}

	userId, err := strconv.Atoi(ctx.GetString("id_user"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": err.Error(),
		})
		return
	}

	merchantData := models.GetMerchantByUserID(userId)
	if merchantData == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Data Merchant tidak ditemukan.",
		})
		return
	}

	qrisData, errQRis := library.Decode(merchantData.QrisData)
	if errQRis != nil {
		logger.Log.Err(errQRis).Msg("Error Body Generate QRIS:")
		ctx.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Gagal membuat QRIS, mohon hubungi admin.",
		})
		return
	}

	type ResponseQris struct {
		QRData       string `json:"qr_data"`
		NamaMerchant string `json:"nama_merchant"`
		MNIDMerchant string `json:"mnid_merchant"`
	}

	result := ResponseQris{
		QRData:       library.ConvertQris(qrisData, int64(trx_model.Amount)),
		NamaMerchant: merchantData.QrisName,
		MNIDMerchant: merchantData.QrisMnid,
	}

	ctx.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Berhasil Membuat QRIS",
		"data":    result,
	})
}
