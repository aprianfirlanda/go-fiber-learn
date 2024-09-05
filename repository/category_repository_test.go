package repository

import (
	"github.com/stretchr/testify/assert"
	"go-fiber-learn/app"
	"go-fiber-learn/model/domain"
	"testing"
)

func TestCategoryRepositoryImpl_Create(t *testing.T) {
	app.LoadEnvironment("..")
	db := app.OpenConnection()
	repository := NewCategoryRepository()
	category := &domain.Category{
		Name: "spend",
	}
	err := repository.Create(db, category)
	assert.Nil(t, err)
	assert.NotEqual(t, 0, category.ID)
}
