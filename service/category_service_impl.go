package service

import (
	"github.com/sirupsen/logrus"
	"go-fiber-learn/model/domain"
	"go-fiber-learn/model/web"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB *gorm.DB
}

func NewCategoryServiceImpl(db *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{DB: db}
}

func (service *CategoryServiceImpl) Create(request *web.CategoryCreateRequest) {
	category := domain.Category{Name: request.Name}
	err := service.DB.Create(&category).Error
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}
