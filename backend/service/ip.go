package service

import (
	"toychart/model"
	"toychart/repository"
)

type IPService struct {
	ipRepo repository.IPRepository
}

func NewIPService(ipRepo repository.IPRepository) *IPService {
	return &IPService{ipRepo: ipRepo}
}

func (s *IPService) Create(ip *model.IP) error {
	return s.ipRepo.Create(ip)
}

func (s *IPService) GetById(id string) (*model.IP, error) {
	return s.ipRepo.GetById(id)
}

func (s *IPService) GetByIPName(ipName string) (*model.IP, error) {
	return s.ipRepo.GetByIPName(ipName)
}

func (s *IPService) CheckIPExists(ipName string) (bool, error) {
	return s.ipRepo.CheckIPExists(ipName)
}

func (s *IPService) GetAllIPsByIPTypes(ipTypeId string) ([]*model.IP, error) {
	return s.ipRepo.GetAllIPsByIPTypes(ipTypeId)
}

func (s *IPService) Update(ip *model.IP) error {
	return s.ipRepo.Update(ip)
}

func (s *IPService) Delete(id string) error {
	return s.ipRepo.Delete(id)
}
