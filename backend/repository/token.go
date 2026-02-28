package repository

import (
	"errors"
	"time"
	"toychart/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type TokenRepository interface {
	Create(user *model.Token) error
	GetById(id string) (*model.Token, error)
	FindByReferenceIdAndDeviceId(referenceId, deviceId string) (*model.Token, error)
	FindByReferenceIdAndToken(referenceId, accessToken string) (*model.Token, error)
	GetByDeviceId(deviceId string) (*model.Token, error)
	Upsert(token *model.Token) error
	Update(user *model.Token) error
	Delete(id string) error
}

type tokenRepository struct {
	db *gorm.DB
}

func NewTokenRepository(db *gorm.DB) TokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(token *model.Token) error {
	result := r.db.Create(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *tokenRepository) GetById(id string) (*model.Token, error) {
	var user model.Token
	result := r.db.First(&user, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &user, nil
}

func (r *tokenRepository) FindByReferenceIdAndDeviceId(referenceId, deviceId string) (*model.Token, error) {
	var token model.Token
	result := r.db.
		Where("referenceId = ? AND deviceId = ?", referenceId, deviceId).
		First(&token)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &token, nil
}

func (r *tokenRepository) FindByReferenceIdAndToken(referenceId, accessToken string) (*model.Token, error) {
	var token model.Token
	result := r.db.
		Where("referenceId = ? AND accessToken = ?", referenceId, accessToken).
		First(&token)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &token, nil
}

func (r *tokenRepository) GetByDeviceId(deviceId string) (*model.Token, error) {
	var token model.Token
	result := r.db.
		Where("deviceId = ?", deviceId).
		First(&token)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &token, nil
}

func (r *tokenRepository) Upsert(token *model.Token) error {
	result := r.db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(token)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *tokenRepository) Update(token *model.Token) error {
	result := r.db.Save(token)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *tokenRepository) Delete(id string) error {
	result := r.db.Model(&model.Token{}).Where("id = ?", id).Updates(map[string]any{
		"status":          false,
		"updatedDateTime": time.Now().UTC(),
	})
	return result.Error
}
