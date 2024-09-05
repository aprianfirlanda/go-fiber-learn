package service

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"go-fiber-learn/app"
	"go-fiber-learn/model/domain"
	"go-fiber-learn/model/web"
	"go-fiber-learn/repository"
	"testing"
)

func TestCategoryServiceImpl_Create(t *testing.T) {
	app.LoadEnvironment("test")
	db := app.OpenConnection()
	categoryRepositoryMock := new(repository.CategoryRepositoryMock)
	categoryRepositoryMock.Mock.On("Create", db, &domain.Category{Name: "spend"}).Return(nil)
	service := NewCategoryService(db, categoryRepositoryMock)
	categoryId, err := service.Create(&web.CategoryCreateRequest{Name: "spend"})
	assert.Nil(t, err, "Create Category return an error")
	logrus.Info("categoryId :" + string(categoryId))
	assert.NotEqual(t, 0, categoryId)
}
