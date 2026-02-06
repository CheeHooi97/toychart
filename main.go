package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"toychart/config"
	"toychart/database"
	"toychart/handler"
	"toychart/repository"
	"toychart/router"
	"toychart/service"

	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	// Load config
	config.LoadConfig()

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Kuala_Lumpur",
		config.DBHost,
		config.DBUser,
		config.DBPassword,
		config.DBName,
		config.DBPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "toychart.",
			SingularTable: true,
			NoLowerCase:   false,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate
	if err := database.Migrate(db); err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	// Init layers
	repos := repository.InitializeRepository(db)
	services := service.InitializeService(repos)
	h := handler.NewHandler(services)

	api := router.SetupRoutes(h, db)

	e := echo.New()
	e.HTTPErrorHandler = func(err error, c echo.Context) {
		if !c.Response().Committed {
			c.JSON(http.StatusInternalServerError, map[string]any{
				"error": map[string]any{
					"code":    "INTERNAL_ERROR",
					"message": "Internal error",
					"debug":   err.Error(),
				},
			})
		}
	}

	e.Any("/*", func(c echo.Context) error {
		api.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	go func() {
		if err := e.Start(":8080"); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
