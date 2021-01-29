package categoryService

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
)

type CategoryServiceInterface interface {
	GetCategory(id uint) *categoryRepository.Category
	GetCategories(params *requestService.RequestParams) []categoryRepository.Category
	CreateCategory(params *requestService.RequestParams) (*categoryRepository.Category, error)
	UpdateCategory(params *requestService.RequestParams, category *categoryRepository.Category) error
	DeleteCategory(category *categoryRepository.Category)
}