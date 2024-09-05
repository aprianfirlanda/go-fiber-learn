package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/app"
	"go-fiber-learn/handler"
	"go-fiber-learn/service"
	"time"
)

func main() {
	fiberApp := fiber.New(fiber.Config{
		IdleTimeout:  time.Second * 5,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	})

	app.LoadEnvironment("run")
	db := app.OpenConnection()
	categoryService := service.NewCategoryService(db)
	handler.CategoryHandler(fiberApp, categoryService)

	err := fiberApp.Listen("localhost:3000")
	if err != nil {
		panic(err)
	}
}
