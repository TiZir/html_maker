package main

import (
	"log"
	"os"

	"github.com/TiZir/html_maker/api"
	"github.com/TiZir/html_maker/db"
	"github.com/TiZir/html_maker/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func run() error {
	e := echo.New()
	database, err := db.GetDB()
	if err != nil {
		log.Printf("Error connect db: %v", err)
	}
	defer database.Close()
	repository := db.NewRepository(database)
	service := api.NewAPI(repository)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", handler.HomePage)
	e.GET("/books", func(c echo.Context) error {
		return handler.GetBooks(c, service)
	})

	// Start server
	e.Logger.Fatal(e.Start(os.Getenv("HTTP_HOST") + ":" + os.Getenv("HTTP_PORT")))
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
