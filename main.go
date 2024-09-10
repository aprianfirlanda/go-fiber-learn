package main

import (
	"errors"
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/app"
	"go-fiber-learn/exception"
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
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			var sqlError *exception.SqlError
			if errors.As(err, &sqlError) {
				return ctx.
					Status(fiber.StatusInternalServerError).
					JSON(fiber.Map{
						"message": sqlError.Message,
					})
			}

			return ctx.
				Status(fiber.StatusInternalServerError).
				JSON(fiber.Map{
					"message": err.Error(),
				})
		},
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
