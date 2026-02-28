package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type UserToyRepository interface {
	Create(auth *model.UserToy) error
	GetById(id string) (*model.UserToy, error)
	GetAllToys() ([]*model.UserToy, error)
	Update(auth *model.UserToy) error
	Delete(id string) error
}

type userToyRepository struct {
	db *gorm.DB
}

func NewUserToyRepository(db *gorm.DB) UserToyRepository {
	return &userToyRepository{db: db}
}

func (r *userToyRepository) Create(auth *model.UserToy) error {
	result := r.db.Create(auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userToyRepository) GetById(id string) (*model.UserToy, error) {
	var auth model.UserToy
	result := r.db.First(&auth, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &auth, nil
}

func (r *userToyRepository) GetAllToys() ([]*model.UserToy, error) {
	var UserToy []*model.UserToy
	result := r.db.
		Order("createdDateTime ASC").
		Find(&UserToy)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return UserToy, nil
}

func (r *userToyRepository) Update(auth *model.UserToy) error {
	result := r.db.Save(auth)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *userToyRepository) Delete(id string) error {
	result := r.db.Model(&model.UserToy{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
