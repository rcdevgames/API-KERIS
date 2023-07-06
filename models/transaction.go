package models

import (
	"QAPI/configs"
	"QAPI/entities"
	"QAPI/logger"
)

func GetMerchantCheckMutation() (merchants []entities.Merchant) {
	err := configs.DB.Unscoped().
		Where("last_mutation_found=?", 0).
		Order("last_mutation_date asc").
		Find(&merchants).Error
	if err != nil {
		logger.Log.Err(err).Msg("Error Get list merchant:")
	}
	return
}
