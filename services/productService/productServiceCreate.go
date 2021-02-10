package productService

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

//сервис создрания товара
func (s *productService) CreateProduct(params *requestService.RequestParams) (*productRepository.Product, error) {
	product := &productRepository.Product{}

	for _, proc := range s.processorsValidate {
		if !proc.IsValid(params) {
			return nil, proc.GetError()
		}
	}

	product.Name = params.Name
	product.SKU = params.SKU
	product.CategoryID = params.CategoryID

	return product, nil
}