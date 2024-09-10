package service

import (
	"go-fiber-learn/exception"
	"go-fiber-learn/model/domain"
	"go-fiber-learn/model/web"
	"go-fiber-learn/repository"
	"gorm.io/gorm"
)

type CategoryServiceImpl struct {
	DB                 *gorm.DB
	CategoryRepository repository.CategoryRepository
}

func NewCategoryService(db *gorm.DB, categoryRepository repository.CategoryRepository) CategoryService {
	return &CategoryServiceImpl{
		DB:                 db,
		CategoryRepository: categoryRepository,
	}
}

func (service *CategoryServiceImpl) Create(request *web.CategoryCreateRequest) (int32, error) {
	category := &domain.Category{Name: request.Name}
	err := service.CategoryRepository.Create(service.DB, category)
	if err != nil {
		return 0, &exception.SqlError{Message: err.Error()}
	}
	return category.ID, nil
}
