package service

import (
	"toychart/model"
	"toychart/repository"
)

type SetService struct {
	setRepo repository.SetRepository
}

func NewSetService(setRepo repository.SetRepository) *SetService {
	return &SetService{setRepo: setRepo}
}

func (s *SetService) Create(set *model.Set) error {
	return s.setRepo.Create(set)
}

func (s *SetService) GetById(id string) (*model.Set, error) {
	return s.setRepo.GetById(id)
}

func (s *SetService) GetBySeries(series string) ([]*model.Set, error) {
	return s.setRepo.GetBySeries(series)
}

func (s *SetService) GetSeriesByIPName(set string) ([]*model.Set, error) {
	return s.setRepo.GetSeriesByIPName(set)
}

func (s *SetService) Update(set *model.Set) error {
	return s.setRepo.Update(set)
}

func (s *SetService) Delete(id string) error {
	return s.setRepo.Delete(id)
}
