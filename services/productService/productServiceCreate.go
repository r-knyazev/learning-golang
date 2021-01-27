package productService

import (
	"errors"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"log"
)

//сервис создрания товара
func (s *productService) CreateProduct(params *requestService.RequestParams) (*productRepository.Product, error) {
	product := &productRepository.Product{}

	for _, proc := range s.processorsValidate {
		if !proc.IsValid(params) {
			return product, proc.GetError()
		}
	}

	product.Name = params.Name
	product.SKU = params.SKU
	product.CategoryID = params.CategoryID

	product, err := productRepository.Repository.Save(product)
	if err != nil {
		log.Println(err)

		return product, errors.New("error while save product")
	}

	return product, nil
}