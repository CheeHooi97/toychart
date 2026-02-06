package model

import (
	"time"
	"toychart/utils"
)

type UserToy struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	UserId   string `gorm:"type:varchar(255)" json:"userId"`
	ToyId    string `gorm:"type:varchar(255)" json:"toyId"`
	Quantity int64  `gorm:"type:bigint" json:"quantity"`
	BaseModel
}

func NewUserToy() *UserToy {
	now := time.Now().UTC()

	m := new(UserToy)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *UserToy) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *UserToy) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
