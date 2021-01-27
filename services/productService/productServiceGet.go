package productService

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

//получение товара по ID
//если товар не найден, вернет nil
func (s *productService) GetProduct(id uint) *productRepository.Product {
	return productRepository.Repository.GetById(id)
}

//получение товаров по фильтрам
func (s *productService) GetProducts(params *requestService.RequestParams) []productRepository.Product {
	var conditions productRepository.FindByConditions

	for _, proc := range s.processorsCond {
		proc.GetCond(params, &conditions)
	}

	return productRepository.Repository.FindBy(conditions)
}