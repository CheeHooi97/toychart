package model

import (
	"time"
	"toychart/utils"
)

type ToyPrice struct {
	Id    string  `gorm:"type:varchar(36);primaryKey" json:"id"`
	ToyId string  `gorm:"type:varchar(255)" json:"toyId"`
	Price float64 `gorm:"type:double precision;default:0" json:"price"`
	BaseModel
}

func NewToyPrice() *ToyPrice {
	now := time.Now().UTC()

	m := new(ToyPrice)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *ToyPrice) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *ToyPrice) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
