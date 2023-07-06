package models

import (
	"QAPI/configs"
	"QAPI/entities"
	"QAPI/library"
	"time"
)

func GetUserByIdDevice(deviceId string) (users *entities.Users) {
	err := configs.DB.Where("device_id = ?", deviceId).First(&users).Error
	if err != nil {
		library.Log.Err(err).Msg("Error Get User By DeviceID:")
		users = nil
	}
	return
}

func CreateUser(data entities.UserInsert) (status bool) {
	err := configs.DB.Create(data).Error
	if err != nil {
		library.Log.Err(err).Msg("Error Create User:")
		status = false
		return
	}
	status = true
	return
}

func UpdateLastLogin(id int8) (status bool) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	users := map[string]interface{}{
		"last_login": time.Now().In(loc),
	}
	err := configs.DB.Model(&entities.Users{}).Where("id=?", id).Updates(users).Error
	if err != nil {
		library.Log.Err(err).Msg("Error Update Last Login user:")
		return false
	}
	return true
}
