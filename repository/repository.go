package repository

import "gorm.io/gorm"

type Repositories struct {
	ToyRepo              ToyRepository
	ToyPriceRepo         ToyPriceRepository
	SetRepo              SetRepository
	UserRepo             UserRepository
	UserToyRepo          UserToyRepository
	UserToySearchLogRepo UserToySearchLogRepository
}

func InitializeRepository(db *gorm.DB) *Repositories {
	return &Repositories{
		ToyRepo:              NewToyRepository(db),
		ToyPriceRepo:         NewToyPriceRepository(db),
		SetRepo:              NewSetRepository(db),
		UserToySearchLogRepo: NewUserToySearchLogRepository(db),
		UserToyRepo:          NewUserToyRepository(db),
		UserRepo:             NewUserRepository(db),
	}
}
