package model

import (
	"time"
	"toychart/utils"
)

type Toy struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	Name     string `gorm:"type:varchar(255)" json:"name"`
	SetName  string `gorm:"type:varchar(255)" json:"setName"`
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewToy() *Toy {
	now := time.Now().UTC()

	m := new(Toy)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *Toy) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *Toy) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
