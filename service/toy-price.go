package service

import (
	"toychart/model"
	"toychart/repository"
)

type ToyPriceService struct {
	ToyPriceRepo repository.ToyPriceRepository
}

func NewToyPriceService(ToyPriceRepo repository.ToyPriceRepository) *ToyPriceService {
	return &ToyPriceService{ToyPriceRepo: ToyPriceRepo}
}

func (s *ToyPriceService) Create(ToyPrice *model.ToyPrice) error {
	return s.ToyPriceRepo.Create(ToyPrice)
}

func (s *ToyPriceService) GetById(id string) (*model.ToyPrice, error) {
	return s.ToyPriceRepo.GetById(id)
}

func (s *ToyPriceService) GetByToyPriceNameAndSet(ToyPriceName, set string) bool {
	return s.ToyPriceRepo.GetByToyPriceNameAndSet(ToyPriceName, set)
}

func (s *ToyPriceService) Update(ToyPrice *model.ToyPrice) error {
	return s.ToyPriceRepo.Update(ToyPrice)
}

func (s *ToyPriceService) Delete(id string) error {
	return s.ToyPriceRepo.Delete(id)
}
