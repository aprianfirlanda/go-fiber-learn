package service

import "go-fiber-learn/model/domain"

type CategoryService interface {
	Create(category domain.Category)
}
