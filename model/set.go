package model

import (
	"time"
	"toychart/utils"
)

type Set struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	SetName  string `gorm:"type:varchar(255)" json:"setName"`
	Series   string `gorm:"type:varchar(255)" json:"series"`
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewSet() *Set {
	now := time.Now().UTC()

	m := new(Set)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *Set) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *Set) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
