package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"fmt"

	"github.com/konnenl/learning-system/config"
	"github.com/konnenl/learning-system/internal/handler"
	"github.com/konnenl/learning-system/internal/database"
	"github.com/konnenl/learning-system/internal/service"
)

func main() {
	fmt.Println("dneo")
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %q", err)
	}
	db, err := database.New()
	if err != nil {
		log.Fatalf("Failed to initialize database: %q", err)
	}
	err = database.Migrate(db)
	if err != nil {
		log.Fatalf("Migration failed: %v", err)
	}
	log.Println("Migration completed successfully!")
	defer func() {
		if sqlDB, err := db.DB(); err == nil {
			sqlDB.Close()
		}
	}()

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${error}` + "\n",
	}))

	service := service.NewService()

	handlers := handler.NewHandler(service)
	handlers.InitRoutes(e)

	port := ":" + cfg.ServerPort
	e.Logger.Fatal(e.Start(port))
}
