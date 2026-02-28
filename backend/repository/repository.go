package repository

import "gorm.io/gorm"

type Repositories struct {
	TokenRepo            TokenRepository
	ToyRepo              ToyRepository
	ToyPriceRepo         ToyPriceRepository
	SetRepo              SetRepository
	UserRepo             UserRepository
	UserDeviceRepo       UserDeviceRepository
	UserToyRepo          UserToyRepository
	UserToySearchLogRepo UserToySearchLogRepository
}

func InitializeRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		TokenRepo:            NewTokenRepository(db),
		ToyRepo:              NewToyRepository(db),
		ToyPriceRepo:         NewToyPriceRepository(db),
		SetRepo:              NewSetRepository(db),
		UserRepo:             NewUserRepository(db),
		UserDeviceRepo:       NewUserDeviceRepository(db),
		UserToySearchLogRepo: NewUserToySearchLogRepository(db),
		UserToyRepo:          NewUserToyRepository(db),
	}
}
