package model

import (
	"time"
	"toychart/utils"
)

type UserToySearchLog struct {
	Id    string `gorm:"type:varchar(36);primaryKey" json:"id"`
	ToyId string `gorm:"type:varchar(255)" json:"toyId"`
	BaseModel
}

func NewUserToySearchLog() *UserToySearchLog {
	now := time.Now().UTC()

	m := new(UserToySearchLog)
	m.Id = utils.UniqueID()
	m.CreatedDateTime = now
	m.UpdatedDateTime = now

	return m
}

func (m *UserToySearchLog) DateTime() {
	m.CreatedDateTime = time.Now().UTC()
	m.UpdatedDateTime = time.Now().UTC()
}

func (m *UserToySearchLog) UpdateDt() {
	m.UpdatedDateTime = time.Now().UTC()
}
