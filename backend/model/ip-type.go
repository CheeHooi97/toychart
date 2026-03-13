package model

import (
	"time"
	"toychart/utils"
)

type IPType struct {
	Id         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	IPTypeName string `gorm:"type:varchar(64);index" json:"ipTypeName"` //popmart, funco etc
	PhotoUrl   string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewIPType() *IPType {
	now := time.Now().UTC()

	m := new(IPType)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *IPType) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *IPType) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
