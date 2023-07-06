package models

import (
	"QAPI/configs"
	"QAPI/entities"
	"QAPI/logger"
	"time"
)

func InsGetUser(deviceId string) (users *entities.Users) {
	input := entities.Users{
		DeviceId:  deviceId,
		CreatedBy: deviceId,
	}
	if err := configs.DB.Where("device_id = ?", deviceId).Assign(input).FirstOrCreate(&users).Error; err != nil {
		logger.Log.Err(err).Msg("Error Get User By DeviceID:")
		users = nil
	}
	return
}

func GetUserByIdDevice(deviceId string) (users *entities.Users) {
	err := configs.DB.Where("device_id = ?", deviceId).First(&users).Error
	if err != nil {
		logger.Log.Err(err).Msg("Error Get User By DeviceID:")
		users = nil
	}
	return
}

func CreateUser(data entities.UserInsert) (status bool) {
	err := configs.DB.Exec("INSERT INTO users(device_id, onesignal_id, created_by) values (?,?,?)", data.DeviceId, data.OnesignalId, data.DeviceId).Error
	if err != nil {
		logger.Log.Err(err).Msg("Error Create User:")
		status = false
		return
	}
	status = true
	return
}

func UpdateLastLogin(id int) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	users := map[string]interface{}{
		"last_login": time.Now().In(loc),
	}
	err := configs.DB.Model(&entities.Users{}).Where("id=?", id).Updates(users).Error
	if err != nil {
		logger.Log.Err(err).Msg("Error Update Last Login user:")
	}
}
