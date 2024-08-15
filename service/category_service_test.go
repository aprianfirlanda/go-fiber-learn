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
	err := serviceImpl.Create(category)
	assert.Nil(t, err)
}

func TestCategoryServiceImplError(t *testing.T) {
	serviceImpl := NewCategoryServiceImpl(db)
	err := serviceImpl.Create(category)
	assert.NotNil(t, err)
}
