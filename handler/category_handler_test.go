package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
	"go-fiber-learn/model/web"
	"go-fiber-learn/service"
	"io"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCategoryHandler(t *testing.T) {
	categoryServiceMock := new(service.CategoryServiceMock)
	categoryServiceMock.Mock.On("Create", &web.CategoryCreateRequest{Name: "spend"}).Return(int32(1), nil)
	fiberApp := fiber.New()
	CategoryHandler(fiberApp, categoryServiceMock)

	reqBody := strings.NewReader(`{"name": "spend"}`)
	req := httptest.NewRequest("POST", "/api/v1/category/new", reqBody)
	req.Header.Set("Content-Type", "application/json")
	res, err := fiberApp.Test(req)
	assert.Nil(t, err)
	assert.Equal(t, 201, res.StatusCode)
	bytes, err := io.ReadAll(res.Body)
	assert.Nil(t, err)
	assert.Equal(t, `{"category_id":1}`, string(bytes))
}
