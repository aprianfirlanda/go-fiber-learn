package service

import (
	"go-fiber-learn/model/web"
)

type CategoryService interface {
	Create(request *web.CategoryCreateRequest)
}
