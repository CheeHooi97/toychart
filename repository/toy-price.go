package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type ToyPriceRepository interface {
	Create(ToyPrice *model.ToyPrice) error
	GetById(id string) (*model.ToyPrice, error)
	GetByToyPriceNameAndSet(ToyPriceName, set string) bool
	Update(ToyPrice *model.ToyPrice) error
	Delete(id string) error
}

type toyPriceRepository struct {
	db *gorm.DB
}

func NewToyPriceRepository(db *gorm.DB) ToyPriceRepository {
	return &toyPriceRepository{db: db}
}

func (r *toyPriceRepository) Create(ToyPrice *model.ToyPrice) error {
	result := r.db.Create(ToyPrice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *toyPriceRepository) GetById(id string) (*model.ToyPrice, error) {
	var ToyPrice model.ToyPrice
	result := r.db.First(&ToyPrice, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ToyPrice, nil
}

func (r *toyPriceRepository) GetByToyPriceNameAndSet(ToyPriceName, set string) bool {
	var auth model.ToyPrice
	result := r.db.
		Where("name = ? AND set = ?", ToyPriceName, set).
		First(&auth)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false
		}
	}
	return true
}

func (r *toyPriceRepository) Update(ToyPrice *model.ToyPrice) error {
	result := r.db.Save(ToyPrice)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *toyPriceRepository) Delete(id string) error {
	result := r.db.Model(&model.ToyPrice{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
