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
	app.LoadEnvironment()
	db = app.OpenConnection()

	m.Run()

	db.Delete(&domain.Category{}).Where("Name = 'satu'")
}

func TestCategoryServiceImplCreate(t *testing.T) {
	serviceImpl := NewCategoryServiceImpl(db)
	assert.NotPanics(t, func() {
		serviceImpl.Create(&web.CategoryCreateRequest{Name: "satu"})
	}, "The code did panic as not expected")
}
