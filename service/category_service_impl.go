package service

import (
	"github.com/sirupsen/logrus"
	"go-fiber-learn/model/domain"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	db *gorm.DB
}

func NewCategoryServiceImpl(db *gorm.DB) *CategoryServiceImpl {
	return &CategoryServiceImpl{db: db}
}

func (service CategoryServiceImpl) Create(category *domain.Category) {
	err := service.db.Create(&category).Error
	if err != nil {
		logrus.Error(err)
		panic(err)
	}
}
