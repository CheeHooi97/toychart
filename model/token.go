package model

import (
	"time"
	"toychart/utils"
)

type Token struct {
	Id          string `gorm:"type:varchar(36);primaryKey" json:"id"`
	ReferenceId string `gorm:"type:varchar(255)" json:"referenceId"`
	DeviceId    string `gorm:"type:varchar(255)" json:"deviceId"`
	AccessToken string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewToken() *Token {
	now := time.Now().UTC()

	m := new(Token)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *Token) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}

type UserWithToken struct {
	*User
	Token string `json:"token"`
}
