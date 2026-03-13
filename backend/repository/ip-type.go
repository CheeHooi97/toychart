package repository

import (
	"errors"
	"toychart/model"

	"gorm.io/gorm"
)

type IPTypeRepository interface {
	Create(IPType *model.IPType) error
	GetById(id string) (*model.IPType, error)
	GetByName(name string) (*model.IPType, error)
	CheckIPTypeExists(ipName string) (bool, error)
	GetAllIPTypes() ([]*model.IPType, error)
	Update(IPType *model.IPType) error
	Delete(id string) error
}

type ipTypeRepository struct {
	db *gorm.DB
}

func NewIPTypeRepository(db *gorm.DB) IPTypeRepository {
	return &ipTypeRepository{db: db}
}

func (r *ipTypeRepository) Create(IPType *model.IPType) error {
	result := r.db.Create(IPType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipTypeRepository) GetById(id string) (*model.IPType, error) {
	var ipType model.IPType
	result := r.db.First(&ipType, id)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ipType, nil
}

func (r *ipTypeRepository) GetByName(name string) (*model.IPType, error) {
	var ipType model.IPType
	result := r.db.Where("ip_type_name = ?", name).First(&ipType)
	if result.Error != nil {
		if !errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, result.Error
		}
	}
	return &ipType, nil
}

func (r *ipTypeRepository) CheckIPTypeExists(ipName string) (bool, error) {
	var ipType model.IPType
	result := r.db.Where("ip_type_name = ?", ipName).First(&ipType)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
	}
	return true, nil
}

func (r *ipTypeRepository) GetAllIPTypes() ([]*model.IPType, error) {
	var ipTypes []*model.IPType
	result := r.db.Find(&ipTypes)
	if result.Error != nil {
		return nil, result.Error
	}
	return ipTypes, nil
}

func (r *ipTypeRepository) Update(IPType *model.IPType) error {
	result := r.db.Save(IPType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *ipTypeRepository) Delete(id string) error {
	result := r.db.Model(&model.IPType{}).Where("id = ?", id).Update("status", false)
	return result.Error
}
