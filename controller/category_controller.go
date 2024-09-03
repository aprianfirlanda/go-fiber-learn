package controller

import "github.com/gofiber/fiber/v2"

type CategoryController interface {
	Create(ctx *fiber.Ctx) error
}
