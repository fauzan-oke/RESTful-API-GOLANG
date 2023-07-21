package repository

import "github.com/fauzan-oke/RESTful-API-GOLANG/entity"

type CategoryRepository interface {
	FindById(id string) *entity.Category
}
