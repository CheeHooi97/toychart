package service

import "toychart/repository"

type Services struct {
	ToyService              *ToyService
	ToyPriceService         *ToyPriceService
	SetService              *SetService
	UserService             *UserService
	UserToyService          *UserToyService
	UserToySearchLogService *UserToySearchLogService
}

func InitializeService(repos *repository.Repositories) *Services {
	return &Services{
		ToyService:              NewToyService(repos.ToyRepo),
		ToyPriceService:         NewToyPriceService(repos.ToyPriceRepo),
		SetService:              NewSetService(repos.SetRepo),
		UserService:             NewUserService(repos.UserRepo),
		UserToyService:          NewUserToyService(repos.UserToyRepo),
		UserToySearchLogService: NewUserToySearchLogService(repos.UserToySearchLogRepo),
	}
}
