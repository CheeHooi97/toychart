package service

import (
	"toychart/model"
	"toychart/repository"
)

type IPSeriesItemService struct {
	ipSeriesItemRepo repository.IPSeriesItemRepository
}

func NewIPSeriesItemService(ipSeriesItemRepo repository.IPSeriesItemRepository) *IPSeriesItemService {
	return &IPSeriesItemService{ipSeriesItemRepo: ipSeriesItemRepo}
}

func (s *IPSeriesItemService) Create(IPSeriesItem *model.IPSeriesItem) error {
	return s.ipSeriesItemRepo.Create(IPSeriesItem)
}

func (s *IPSeriesItemService) GetById(id string) (*model.IPSeriesItem, error) {
	return s.ipSeriesItemRepo.GetById(id)
}

func (s *IPSeriesItemService) GetByItem(item string) (*model.IPSeriesItem, error) {
	return s.ipSeriesItemRepo.GetByItem(item)
}

func (s *IPSeriesItemService) CheckIPSeriesItemExists(item string) (bool, error) {
	return s.ipSeriesItemRepo.CheckIPSeriesItemExists(item)
}

func (s *IPSeriesItemService) GetByIPSeriesId(ipSeriesId string) ([]*model.IPSeriesItem, error) {
	return s.ipSeriesItemRepo.GetByIPSeriesId(ipSeriesId)
}

func (s *IPSeriesItemService) Update(IPSeriesItem *model.IPSeriesItem) error {
	return s.ipSeriesItemRepo.Update(IPSeriesItem)
}

func (s *IPSeriesItemService) Delete(id string) error {
	return s.ipSeriesItemRepo.Delete(id)
}
