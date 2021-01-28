package productService

import (
	"errors"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"log"
)

//сервис обновления товара товара
func (s *productService) UpdateProduct(params *requestService.RequestParams, product *productRepository.Product) error {

	for _, proc := range s.processorsValidate {
		if !proc.IsValid(params) {
			return proc.GetError()
		}
	}

	product.Name = params.Name
	product.SKU = params.SKU
	product.CategoryID = params.CategoryID

	product, err := productRepository.Repository.Save(product)

	if err != nil {
		log.Println(err)

		return errors.New("error while save product")
	}

	return nil
}