package service

import (
	"toychart/model"
	"toychart/repository"
)

type TokenService struct {
	tokenRepo repository.TokenRepository
}

func NewTokenService(tokenRepo repository.TokenRepository) *TokenService {
	return &TokenService{tokenRepo: tokenRepo}
}

func (s *TokenService) Create(token *model.Token) error {
	return s.tokenRepo.Create(token)
}

func (s *TokenService) GetById(id string) (*model.Token, error) {
	return s.tokenRepo.GetById(id)
}

func (s *TokenService) FindByReferenceIdAndDeviceId(referenceId, deviceId string) (*model.Token, error) {
	return s.tokenRepo.FindByReferenceIdAndDeviceId(referenceId, deviceId)
}

func (s *TokenService) FindByReferenceIdAndToken(referenceId, accessToken string) (*model.Token, error) {
	return s.tokenRepo.FindByReferenceIdAndToken(referenceId, accessToken)
}

func (s *TokenService) GetByDeviceId(deviceId string) (*model.Token, error) {
	return s.tokenRepo.GetByDeviceId(deviceId)
}

func (s *TokenService) Upsert(token *model.Token) error {
	return s.tokenRepo.Upsert(token)
}

func (s *TokenService) Update(token *model.Token) error {
	return s.tokenRepo.Update(token)
}

func (s *TokenService) Delete(id string) error {
	return s.tokenRepo.Delete(id)
}
