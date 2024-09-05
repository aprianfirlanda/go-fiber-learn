package repository

import (
	"github.com/sirupsen/logrus"
	"go-fiber-learn/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct{}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (CategoryRepositoryImpl) Create(tx *gorm.DB, category *domain.Category) error {
	err := tx.Create(&category).Error
	if err != nil {
		logrus.Error(err)
		return err
	}
	return nil
}
