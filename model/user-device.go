package model

import (
	"time"
	"toychart/utils"
)

type LoginType string

const (
	LoginTypeOtp    LoginType = "OTP"
	LoginTypeSocial LoginType = "SOCIAL"
)

type UserDevicePlatform string

const (
	UserDevicePlatformAndroid UserDevicePlatform = "ANDROID"
	UserDevicePlatformIOS     UserDevicePlatform = "IOS"
	UserDevicePlatformWeb     UserDevicePlatform = "WEB"
)

type UserDevice struct {
	Id         string             `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserId     string             `gorm:"type:varchar(255)" json:"userId"`
	Platform   UserDevicePlatform `gorm:"type:varchar(255)" json:"platform"`
	DeviceId   string             `gorm:"type:varchar(255)" json:"deviceId"`
	DeviceInfo string             `gorm:"type:varchar(255)" json:"deviceInfo"`
	PNSToken   string             `gorm:"type:varchar(255)" json:"pnsToken"`
	IsHuawei   bool               `gorm:"type:boolean;default:false" json:"isHuawei"`
	BaseModel
}

func NewUserDevice() *UserDevice {
	now := time.Now().UTC()

	m := new(UserDevice)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *UserDevice) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
