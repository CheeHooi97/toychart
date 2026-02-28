package service

import "toychart/repository"

type Services struct {
	TokenService            *TokenService
	ToyService              *ToyService
	ToyPriceService         *ToyPriceService
	SetService              *SetService
	UserService             *UserService
	UserDeviceService       *UserDeviceService
	UserToyService          *UserToyService
	UserToySearchLogService *UserToySearchLogService
}

func InitializeService(repos *repository.Repositories) *Services {
	return &Services{
		TokenService:            NewTokenService(repos.TokenRepo),
		ToyService:              NewToyService(repos.ToyRepo),
		ToyPriceService:         NewToyPriceService(repos.ToyPriceRepo),
		SetService:              NewSetService(repos.SetRepo),
		UserService:             NewUserService(repos.UserRepo),
		UserDeviceService:       NewUserDeviceService(repos.UserDeviceRepo),
		UserToyService:          NewUserToyService(repos.UserToyRepo),
		UserToySearchLogService: NewUserToySearchLogService(repos.UserToySearchLogRepo),
	}
}
