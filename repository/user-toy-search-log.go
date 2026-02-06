package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type UserToySearchLogRepository interface {
	Create(auth *model.UserToySearchLog) error
	GetById(id string) (*model.UserToySearchLog, error)
	GetAllToys() ([]*model.UserToySearchLog, error)
	Update(auth *model.UserToySearchLog) error
	Delete(id string) error
}

type userToySearchLogRepository struct {
	db *gorm.DB
}

func NewUserToySearchLogRepository(db *gorm.DB) UserToySearchLogRepository {
	return &userToySearchLogRepository{db: db}
}

func (r *userToySearchLogRepository) Create(auth *model.UserToySearchLog) error {
	result := r.db.Create(auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userToySearchLogRepository) GetById(id string) (*model.UserToySearchLog, error) {
	var auth model.UserToySearchLog
	result := r.db.First(&auth, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &auth, nil
}

func (r *userToySearchLogRepository) GetAllToys() ([]*model.UserToySearchLog, error) {
	var UserToySearchLog []*model.UserToySearchLog
	result := r.db.
		Order("createdDateTime ASC").
		Find(&UserToySearchLog)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return UserToySearchLog, nil
}

func (r *userToySearchLogRepository) Update(auth *model.UserToySearchLog) error {
	result := r.db.Save(auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userToySearchLogRepository) Delete(id string) error {
	result := r.db.Model(&model.UserToySearchLog{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
