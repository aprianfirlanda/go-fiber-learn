package repository

import (
	"github.com/stretchr/testify/mock"
	"go-fiber-learn/model/domain"
	"gorm.io/gorm"
)

type CategoryRepositoryMock struct {
	Mock mock.Mock
}

func (categoryRepositoryMock *CategoryRepositoryMock) Create(tx *gorm.DB, category *domain.Category) error {
	arguments := categoryRepositoryMock.Mock.Called(tx, category)
	return arguments.Error(0)
}
