package models

import (
	"QAPI/configs"
	"QAPI/entities"
	"QAPI/logger"
	"fmt"
)

func GetMerchantByUserID(id int) (merchant *entities.Merchant) {
	if err := configs.DB.Where("user_id = ?", id).First(&merchant).Error; err != nil {
		logger.Log.Err(err).Msg("Error Get User By DeviceID:")
		merchant = nil
	}
	return
}

func InsertMerchant(data entities.MerchantInsert, deviceId string) (status bool) {
	fmt.Println("User ID:", data.UserID)
	q := "INSERT INTO merchants(user_id, qris_data, qris_email, qris_password, qris_mnid, qris_name, created_by) value (?,?,?,?,?,?,?)"
	if err := configs.DB.Exec(q, data.UserID, data.QrisData, data.QrisEmail, data.QrisPassword, data.QrisMnid, data.QrisName, deviceId).Error; err != nil {
		logger.Log.Err(err).Msg("Error Insert Merchant:")
		return false
	}
	return true

}

func UpdateMerchant(data map[string]interface{}) (status bool) {
	err := configs.DB.Model(&entities.Merchant{}).Where("id=?", data["id"]).Updates(data).Error
	if err != nil {
		logger.Log.Err(err).Msg("Error Update Last Login user:")
		return false
	}
	return true
}
