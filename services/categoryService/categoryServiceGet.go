package categoryService

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
)

//получение товара по ID
//если товар не найден, вернет nil
func (s *categoryService) GetCategory(id uint) *categoryRepository.Category {
	return categoryRepository.Repository.GetById(id)
}

//получение товаров по фильтрам
func (s *categoryService) GetCategories(params *requestService.RequestParams) []categoryRepository.Category {
	var conditions categoryRepository.FindByConditions

	for _, proc := range s.processorsCond {
		proc.GetCond(params, &conditions)
	}

	return categoryRepository.Repository.FindBy(conditions)
}