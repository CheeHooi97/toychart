package service

import (
	"toychart/model"
	"toychart/repository"
)

type IPTypeService struct {
	ipTypeRepo repository.IPTypeRepository
}

func NewIPTypeService(ipTypeRepo repository.IPTypeRepository) *IPTypeService {
	return &IPTypeService{ipTypeRepo: ipTypeRepo}
}

func (s *IPTypeService) Create(ipType *model.IPType) error {
	return s.ipTypeRepo.Create(ipType)
}

func (s *IPTypeService) GetById(id string) (*model.IPType, error) {
	return s.ipTypeRepo.GetById(id)
}

func (s *IPTypeService) GetByName(name string) (*model.IPType, error) {
	return s.ipTypeRepo.GetByName(name)
}

func (s *IPTypeService) CheckIPTypeExists(ipName string) (bool, error) {
	return s.ipTypeRepo.CheckIPTypeExists(ipName)
}

func (s *IPTypeService) GetAllIPTypes() ([]*model.IPType, error) {
	return s.ipTypeRepo.GetAllIPTypes()
}

func (s *IPTypeService) Update(ipType *model.IPType) error {
	return s.ipTypeRepo.Update(ipType)
}

func (s *IPTypeService) Delete(id string) error {
	return s.ipTypeRepo.Delete(id)
}
