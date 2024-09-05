package service

import (
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	categoryRepositoryMock.Mock.On("Create", db, &domain.Category{Name: "spend"}).Run(func(args mock.Arguments) {
		category := args.Get(1).(*domain.Category)
		category.ID = 1
	}).Return(nil)
	service := NewCategoryService(db, categoryRepositoryMock)
	categoryId, err := service.Create(&web.CategoryCreateRequest{Name: "spend"})
	assert.Nil(t, err, "Create Category return an error")
	logrus.Infof("categoryId : %d", categoryId)
	assert.NotEqual(t, int32(0), categoryId)
}
