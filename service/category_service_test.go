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
	assert.Nil(t, service.Create(&web.CategoryCreateRequest{Name: "satu"}), "Create Category return an error")
}
