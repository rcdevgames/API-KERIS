package entities

import "time"

type Users struct {
	Id          int       `json:"id" gorm:"type:int;primary_key"`
	DeviceId    string    `json:"device_id"`
	OnesignalId string    `json:"onesignal_id"`
	IsActive    int       `json:"is_active"`
	LastLogin   time.Time `json:"last_login"`
	CreatedDate time.Time `json:"created_date"`
	CreatedBy   string    `json:"created_by"`
}

type UserInsert struct {
	DeviceId    string `json:"device_id"`
	OnesignalId string `json:"onesignal_id"`
	CreatedBy   string `json:"created_by"`
}

type UserUpdate struct {
	LastLogin time.Time `json:"last_login"`
}

type LoginModel struct {
	DeviceId    string `json:"device_id"`
	OnesignalId string `json:"onesignal_id"`
}

type LoginResponse struct {
	AccessToken string `json:"access_token"`
}
