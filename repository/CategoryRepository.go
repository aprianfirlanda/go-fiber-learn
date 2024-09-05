package repository

import (
	"go-fiber-learn/model/domain"
	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(tx *gorm.DB, category *domain.Category) error
}
