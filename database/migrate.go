package database

import (
	"log"
	"toychart/model"

	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Migrate runs all migrations into the "toychart" schema
func Migrate(db *gorm.DB) error {
	// Set schema globally using NamingStrategy
	db.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix:   "toychart.", // schema_name.
		SingularTable: true,        // use singular table names
	}

	// List of models
	models := []any{
		&model.Set{},
		&model.Token{},
		&model.ToyPrice{},
		&model.Toy{},
		&model.User{},
		&model.UserDevice{},
		&model.UserToySearchLog{},
		&model.UserToy{},
	}

	// Run AutoMigrate
	err := db.AutoMigrate(models...)
	if err != nil {
		log.Println("Migration error:", err)
		return err
	}

	log.Println("Migration completed successfully!")
	return nil
}
