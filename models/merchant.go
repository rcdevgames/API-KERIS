package models

import "QAPI/entities"

func GetMerchantByUserID(id int8) (merchant *entities.Merchant) {

	return
}

func InsertMerchant(data entities.MerchantInsert) (status bool) {
	return false
}

func UpdateMerchant(data map[string]interface{}) (status bool) {
	return false
}
