package model

import (
	"time"
	"toychart/utils"
)

type Set struct {
	Id         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	IPType     string `gorm:"type:varchar(64);index" json:"ipType"`  //popmart, funco etc
	IPName     string `gorm:"type:varchar(128)" json:"ipName"`       // popmart: labubu, skullpanda etc
	Series     string `gorm:"type:varchar(255);index" json:"series"` // labubu: why-so-serious-series etc
	PhotoUrl   string `gorm:"type:text" json:"photoUrl"`
	ItemName   string `gorm:"type:varchar(255)" json:"itemName"`
	Rarity     string `gorm:"type:varchar(64)" json:"rarity"`
	AvgPrice   string `gorm:"type:varchar(64)" json:"avgPrice"`
	SoldCount  int    `gorm:"type:int;default:0" json:"soldCount"`
	LastSoldAt string `gorm:"type:varchar(64)" json:"lastSoldAt"`
	Href       string `gorm:"type:text" json:"href"`
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
