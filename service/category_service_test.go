package service

import (
	"github.com/stretchr/testify/assert"
	"go-fiber-learn/app"
	"go-fiber-learn/model/domain"
	"go-fiber-learn/model/web"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

func TestMain(m *testing.M) {
	app.LoadEnvironment("test")
	db = app.OpenConnection()

	m.Run()

	db.Where("name = ?", "satu").Delete(&domain.Category{})
}

func TestCategoryServiceImplCreate(t *testing.T) {
	service := NewCategoryService(db)
	categoryId, err := service.Create(&web.CategoryCreateRequest{Name: "satu"})
	assert.Nil(t, err, "Create Category return an error")
	assert.NotEqual(t, 0, categoryId)
}
