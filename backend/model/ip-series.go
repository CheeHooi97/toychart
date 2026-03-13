package model

import (
	"time"
	"toychart/utils"
)

type IPSeries struct {
	Id       string `gorm:"type:varchar(36);primaryKey" json:"id"`
	IPTypeId string `gorm:"type:varchar(36);index" json:"ipTypeId"`
	IPId     string `gorm:"type:varchar(36);index" json:"ipId"`
	Series   string `gorm:"type:varchar(255)" json:"series"`
	PhotoUrl string `gorm:"type:text" json:"photoUrl"`
	BaseModel
}

func NewIPSeries() *IPSeries {
	now := time.Now().UTC()

	m := new(IPSeries)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *IPSeries) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *IPSeries) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
