package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type IPRepository interface {
	Create(ip *model.IP) error
	GetById(id string) (*model.IP, error)
	GetByIPName(ipName string) (*model.IP, error)
	CheckIPExists(ipName string) (bool, error)
	GetAllIPsByIPTypes(ipTypeId string) ([]*model.IP, error)
	Update(ip *model.IP) error
	Delete(id string) error
}

type ipRepository struct {
	db *gorm.DB
}

func NewIPRepository(db *gorm.DB) IPRepository {
	return &ipRepository{db: db}
}

func (r *ipRepository) Create(ip *model.IP) error {
	result := r.db.Create(ip)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipRepository) GetById(id string) (*model.IP, error) {
	var ip model.IP
	result := r.db.First(&ip, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ip, nil
}

func (r *ipRepository) GetByIPName(ipName string) (*model.IP, error) {
	var ip model.IP
	result := r.db.Where("ip_name = ?", ipName).First(&ip)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ip, nil
}

func (r *ipRepository) CheckIPExists(ipName string) (bool, error) {
	var ip model.IP
	result := r.db.Where("ip_name = ?", ipName).First(&ip)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	return true, nil
}

func (r *ipRepository) GetAllIPsByIPTypes(ipTypeId string) ([]*model.IP, error) {
	var ips []*model.IP
	result := r.db.Where("ip_type_id = ?", ipTypeId).Find(&ips)
	if result.Error != nil {
		return nil, result.Error
	}
	return ips, nil
}

func (r *ipRepository) Update(ip *model.IP) error {
	result := r.db.Save(ip)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipRepository) Delete(id string) error {
	result := r.db.Model(&model.IP{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
