package model

import (
	"time"
	"toychart/utils"
)

type User struct {
	Id        string `gorm:"type:varchar(36);primaryKey" json:"id"`
	CompanyId string `gorm:"type:varchar(255)" json:"companyId"`
	Username  string `gorm:"type:varchar(255)" json:"username"`
	Email     string `gorm:"type:varchar(255)" json:"email"`
	PhotoURL  string `gorm:"type:varchar(255)" json:"photoUrl"`
	FcmToken  string `gorm:"type:varchar(255)" json:"fcmToken"`
	Status    bool   `gorm:"type:boolean;default:false" json:"status"`
	BaseModel
}

func NewUser() *User {
	now := time.Now().UTC()

	m := new(User)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *User) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *User) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
