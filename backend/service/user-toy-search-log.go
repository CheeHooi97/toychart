package service

import (
	"toychart/model"
	"toychart/repository"
)

type UserToySearchLogService struct {
	userToySearchLogRepo repository.UserToySearchLogRepository
}

func NewUserToySearchLogService(userToySearchLogRepo repository.UserToySearchLogRepository) *UserToySearchLogService {
	return &UserToySearchLogService{userToySearchLogRepo: userToySearchLogRepo}
}

func (s *UserToySearchLogService) Create(UserToySearchLog *model.UserToySearchLog) error {
	return s.userToySearchLogRepo.Create(UserToySearchLog)
}

func (s *UserToySearchLogService) GetById(id string) (*model.UserToySearchLog, error) {
	return s.userToySearchLogRepo.GetById(id)
}

func (s *UserToySearchLogService) GetAllToys() ([]*model.UserToySearchLog, error) {
	return s.userToySearchLogRepo.GetAllToys()
}

func (s *UserToySearchLogService) Update(UserToySearchLog *model.UserToySearchLog) error {
	return s.userToySearchLogRepo.Update(UserToySearchLog)
}

func (s *UserToySearchLogService) Delete(id string) error {
	return s.userToySearchLogRepo.Delete(id)
}
