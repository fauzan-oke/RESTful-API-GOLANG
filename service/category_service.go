package service

import (
	"errors"

	"github.com/fauzan-oke/RESTful-API-GOLANG/entity"
	"github.com/fauzan-oke/RESTful-API-GOLANG/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepository
}

func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)

	if category == nil {
		return nil, errors.New("not found")
	} else {
		return category, nil
	}
}
