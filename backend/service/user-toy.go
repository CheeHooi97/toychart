package service

import (
	"toychart/model"
	"toychart/repository"
)

type UserToyService struct {
	userToyRepo repository.UserToyRepository
}

func NewUserToyService(userToyRepo repository.UserToyRepository) *UserToyService {
	return &UserToyService{userToyRepo: userToyRepo}
}

func (s *UserToyService) Create(UserToy *model.UserToy) error {
	return s.userToyRepo.Create(UserToy)
}

func (s *UserToyService) GetById(id string) (*model.UserToy, error) {
	return s.userToyRepo.GetById(id)
}

func (s *UserToyService) GetAllToys() ([]*model.UserToy, error) {
	return s.userToyRepo.GetAllToys()
}

func (s *UserToyService) Update(UserToy *model.UserToy) error {
	return s.userToyRepo.Update(UserToy)
}

func (s *UserToyService) Delete(id string) error {
	return s.userToyRepo.Delete(id)
}
