package service

import "toychart/repository"

type Services struct {
	TokenService            *TokenService
	ToyService              *ToyService
	ToyPriceService         *ToyPriceService
	IPService               *IPService
	IPTypeService           *IPTypeService
	IPSeriesService         *IPSeriesService
	IPSeriesItemService     *IPSeriesItemService
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
		IPService:               NewIPService(repos.IPRepo),
		IPTypeService:           NewIPTypeService(repos.IPTypeRepo),
		IPSeriesService:         NewIPSeriesService(repos.IPSeriesRepo),
		IPSeriesItemService:     NewIPSeriesItemService(repos.IPSeriesItemRepo),
		UserService:             NewUserService(repos.UserRepo),
		UserDeviceService:       NewUserDeviceService(repos.UserDeviceRepo),
		UserToyService:          NewUserToyService(repos.UserToyRepo),
		UserToySearchLogService: NewUserToySearchLogService(repos.UserToySearchLogRepo),
	}
}
