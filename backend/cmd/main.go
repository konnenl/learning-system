package main

import(
	"log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/konnenl/learning-system/internal/handler"
	"github.com/konnenl/learning-system/internal/template"
	"github.com/konnenl/learning-system/config"
)

func main(){
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %q", err)
	}

	e := echo.New()
	renderer := template.InitTemplate()
	e.Renderer = renderer
	e.Static("/static", "ui/static")
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `${time_rfc3339} | ${method} | ${uri} | ${status} | ${latency_human} | ${error}` + "\n",
	}))

	handlers := handler.NewHandler()
	handlers.InitRoutes(e)

	port := ":" + cfg.ServerPort
	e.Logger.Fatal(e.Start(port))
}