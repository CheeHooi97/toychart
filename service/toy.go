package service

import (
	"toychart/model"
	"toychart/repository"
)

type ToyService struct {
	ToyRepo repository.ToyRepository
}

func NewToyService(ToyRepo repository.ToyRepository) *ToyService {
	return &ToyService{ToyRepo: ToyRepo}
}

func (s *ToyService) Create(Toy *model.Toy) error {
	return s.ToyRepo.Create(Toy)
}

func (s *ToyService) GetById(id string) (*model.Toy, error) {
	return s.ToyRepo.GetById(id)
}

func (s *ToyService) GetByToyNameAndSet(ToyName, set string) bool {
	return s.ToyRepo.GetByToyNameAndSet(ToyName, set)
}

func (s *ToyService) SearchToyList(keyword, set, order string) ([]*model.Toy, error) {
	return s.ToyRepo.SearchToyList(keyword, set, order)
}

func (s *ToyService) GetAllToys() ([]*model.Toy, error) {
	return s.ToyRepo.GetAllToys()
}

func (s *ToyService) Update(Toy *model.Toy) error {
	return s.ToyRepo.Update(Toy)
}

func (s *ToyService) Delete(id string) error {
	return s.ToyRepo.Delete(id)
}
