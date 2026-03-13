package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type IPSeriesItemRepository interface {
	Create(IPSeriesItem *model.IPSeriesItem) error
	GetById(id string) (*model.IPSeriesItem, error)
	GetByItem(item string) (*model.IPSeriesItem, error)
	CheckIPSeriesItemExists(item string) (bool, error)
	GetByIPSeriesId(ipSeriesId string) ([]*model.IPSeriesItem, error)
	Update(IPSeriesItem *model.IPSeriesItem) error
	Delete(id string) error
}

type ipSeriesItemRepository struct {
	db *gorm.DB
}

func NewIPSeriesItemRepository(db *gorm.DB) IPSeriesItemRepository {
	return &ipSeriesItemRepository{db: db}
}

func (r *ipSeriesItemRepository) Create(IPSeriesItem *model.IPSeriesItem) error {
	result := r.db.Create(IPSeriesItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipSeriesItemRepository) GetById(id string) (*model.IPSeriesItem, error) {
	var IPSeriesItem model.IPSeriesItem
	result := r.db.First(&IPSeriesItem, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &IPSeriesItem, nil
}

func (r *ipSeriesItemRepository) GetByItem(item string) (*model.IPSeriesItem, error) {
	var ipSeriesItem model.IPSeriesItem
	result := r.db.Where("item_name = ?", item).First(&ipSeriesItem)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ipSeriesItem, nil
}

func (r *ipSeriesItemRepository) CheckIPSeriesItemExists(item string) (bool, error) {
	var ipSeriesItem model.IPSeriesItem
	result := r.db.Where("item_name = ?", item).First(&ipSeriesItem)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	return true, nil
}

func (r *ipSeriesItemRepository) GetByIPSeriesId(ipSeriesId string) ([]*model.IPSeriesItem, error) {
	var ipSeriesItems []*model.IPSeriesItem
	result := r.db.Where("ip_series_id = ?", ipSeriesId).Find(&ipSeriesItems)
	if result.Error != nil {
		return nil, result.Error
	}
	return ipSeriesItems, nil
}

func (r *ipSeriesItemRepository) Update(IPSeriesItem *model.IPSeriesItem) error {
	result := r.db.Save(IPSeriesItem)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipSeriesItemRepository) Delete(id string) error {
	result := r.db.Model(&model.IPSeriesItem{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
