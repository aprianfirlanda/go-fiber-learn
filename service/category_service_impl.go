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

func NewCategoryService(db *gorm.DB) CategoryService {
	return &CategoryServiceImpl{DB: db}
}

func (service *CategoryServiceImpl) Create(request *web.CategoryCreateRequest) (int32, error) {
	category := domain.Category{Name: request.Name}
	err := service.DB.Create(&category).Error
	if err != nil {
		logrus.Error(err)
		return 0, err
	}
	return category.ID, nil
}
