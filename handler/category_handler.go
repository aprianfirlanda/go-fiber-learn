package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/model/web"
	"go-fiber-learn/service"
)

func CategoryController(app *fiber.App, categoryService service.CategoryService) {

	app.Post("/", func(ctx *fiber.Ctx) error {
		request := new(web.CategoryCreateRequest)
		_ = ctx.BodyParser(request)
		err := categoryService.Create(request)
		if err != nil {
			return err
		}
		return nil
	})
}
