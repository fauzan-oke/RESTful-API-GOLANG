package service

import (
	"testing"

	"github.com/fauzan-oke/RESTful-API-GOLANG/entity"
	"github.com/fauzan-oke/RESTful-API-GOLANG/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = &repository.CategoryRepositoryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: categoryRepository}

func TestCategoryService_Get(t *testing.T) {
	//program mock
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.NotNil(t, err)
	assert.Nil(t, category)

}

func TestCategoryService_GetFound(t *testing.T) {
	category := entity.Category{
		Id:   "22",
		Name: "Hp",
	}

	categoryRepository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, category.Id, result.Id)
}
