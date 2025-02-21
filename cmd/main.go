package main

import (
	"log"
	"os"

	"pack-calculator/internal/application"
	"pack-calculator/internal/infra/handlers"
	"pack-calculator/internal/infra/repositories"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	repo := repositories.NewInMemoryResultRepository()

	calculatorService := application.NewCalculatorService(repo)

	handlers.Service = calculatorService

	e.GET("/calculate", handlers.CalculateHandler)
	e.Static("/ui", "ui")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Server is running on port %s", port)
	e.Logger.Fatal(e.Start(":" + port))
}
