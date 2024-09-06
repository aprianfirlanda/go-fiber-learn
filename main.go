package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/app"
	"go-fiber-learn/handler"
	"go-fiber-learn/repository"
	"go-fiber-learn/service"
	"time"
)

func main() {
	// load .env file
	app.LoadEnvironment(".")

	// initialize go fiber
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	// Open Connection To PostgreSQL Database
	db := app.OpenConnection()

	// category
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(db, categoryRepository)
	handler.NewCategoryHandler(fiberApp, categoryService)

	// run go fiber
	err := fiberApp.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
