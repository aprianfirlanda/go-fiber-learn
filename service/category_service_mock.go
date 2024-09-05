package service

import (
	"github.com/stretchr/testify/mock"
	"go-fiber-learn/model/web"
)

type CategoryServiceMock struct {
	Mock mock.Mock
}

func (CategoryServiceMock *CategoryServiceMock) Create(request *web.CategoryCreateRequest) (int32, error) {
	arguments := CategoryServiceMock.Mock.Called(request)
	return arguments.Get(0).(int32), arguments.Error(1)
}
