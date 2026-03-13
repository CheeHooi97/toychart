package model

import (
	"time"
	"toychart/utils"
)

type IPSeriesItem struct {
	Id         string `gorm:"type:varchar(36);primaryKey" json:"id"`
	IPTypeId   string `gorm:"type:varchar(36);index" json:"ipTypeId"`
	IPId       string `gorm:"type:varchar(36);index" json:"ipId"`
	IPSeriesId string `gorm:"type:varchar(255);index" json:"ipSeriesId"`
	ItemName   string `gorm:"type:varchar(255)" json:"itemName"`
	PhotoUrl   string `gorm:"type:text" json:"photoUrl"`
	Rarity     string `gorm:"type:varchar(64)" json:"rarity"`
	AvgPrice   string `gorm:"type:varchar(64)" json:"avgPrice"`
	SoldCount  int    `gorm:"type:int;default:0" json:"soldCount"`
	LastSoldAt string `gorm:"type:varchar(64)" json:"lastSoldAt"`
	Href       string `gorm:"type:text" json:"href"`
	BaseModel
}

func NewIPSeriesItem() *IPSeriesItem {
	now := time.Now().UTC()

	m := new(IPSeriesItem)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *IPSeriesItem) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *IPSeriesItem) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
