package controller

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/model/web"
	"go-fiber-learn/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller CategoryControllerImpl) Create(ctx *fiber.Ctx) error {
	request := new(web.CategoryCreateRequest)
	_ = ctx.BodyParser(request)
	controller.CategoryService.Create(request)
	return nil
}
