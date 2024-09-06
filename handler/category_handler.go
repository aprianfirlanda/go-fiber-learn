package handler

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-learn/model/web"
	"go-fiber-learn/service"
)

func NewCategoryHandler(app *fiber.App, categoryService service.CategoryService) {
	apiV1Category := app.Group("/api/v1/category")

	// create new category
	apiV1Category.Post("/new", func(ctx *fiber.Ctx) error {
		request := new(web.CategoryCreateRequest)
		_ = ctx.BodyParser(request)
		categoryId, err := categoryService.Create(request)
		if err != nil {
			return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
			"category_id": categoryId,
		})
	})
}
