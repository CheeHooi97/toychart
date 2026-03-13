package model

import (
	"time"
	"toychart/utils"
)

type IP struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	IPTypeId string `gorm:"type:varchar(36);index" json:"ipTypeId"`
	IPName   string `gorm:"type:varchar(128)" json:"ipName"` // popmart: labubu, skullpanda etc
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewIP() *IP {
	now := time.Now().UTC()

	m := new(IP)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *IP) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *IP) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
