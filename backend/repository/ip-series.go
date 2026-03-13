package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type IPSeriesRepository interface {
	Create(IPSeries *model.IPSeries) error
	GetById(id string) (*model.IPSeries, error)
	GetBySeries(series string) (*model.IPSeries, error)
	CheckIPSeriesExists(series string) (bool, error)
	GetbyIpId(ipId string) ([]*model.IPSeries, error)
	Update(IPSeries *model.IPSeries) error
	Delete(id string) error
}

type ipSeriesRepository struct {
	db *gorm.DB
}

func NewIPSeriesRepository(db *gorm.DB) IPSeriesRepository {
	return &ipSeriesRepository{db: db}
}

func (r *ipSeriesRepository) Create(IPSeries *model.IPSeries) error {
	result := r.db.Create(IPSeries)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipSeriesRepository) GetById(id string) (*model.IPSeries, error) {
	var IPSeries model.IPSeries
	result := r.db.First(&IPSeries, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &IPSeries, nil
}

func (r *ipSeriesRepository) GetBySeries(series string) (*model.IPSeries, error) {
	var IPSeries model.IPSeries
	result := r.db.Where("series = ?", series).First(&IPSeries)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &IPSeries, nil
}

func (r *ipSeriesRepository) CheckIPSeriesExists(series string) (bool, error) {
	var ipSeries model.IPSeries
	result := r.db.Where("series = ?", series).
		First(&ipSeries)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	return true, nil
}

func (r *ipSeriesRepository) GetbyIpId(ipId string) ([]*model.IPSeries, error) {
	var ipSeriesList []*model.IPSeries
	result := r.db.Where("ip_id = ?", ipId).Find(&ipSeriesList)
	if result.Error != nil {
		return nil, result.Error
	}
	return ipSeriesList, nil
}

func (r *ipSeriesRepository) Update(IPSeries *model.IPSeries) error {
	result := r.db.Save(IPSeries)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipSeriesRepository) Delete(id string) error {
	result := r.db.Model(&model.IPSeries{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
