package service

import (
	"github.com/stretchr/testify/assert"
	"go-fiber-learn/app"
	"go-fiber-learn/model/domain"
	"gorm.io/gorm"
	"testing"
)

var db *gorm.DB

var category = &domain.Category{
	ID:   1,
	Name: "Satu",
}

func TestMain(m *testing.M) {
	app.LoadEnvironment()
	db = app.OpenConnection()

	m.Run()

	db.Delete(category)
}

func TestCategoryServiceImplCreate(t *testing.T) {
	serviceImpl := NewCategoryServiceImpl(db)
	assert.NotPanics(t, func() {
		serviceImpl.Create(category)
	}, "The code did panic as not expected")
}

func TestCategoryServiceImplError(t *testing.T) {
	serviceImpl := NewCategoryServiceImpl(db)
	assert.Panics(t, func() {
		serviceImpl.Create(category)
	}, "The code did not panic as expected")
}
