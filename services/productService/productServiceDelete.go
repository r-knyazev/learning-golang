package productService

import (
	"learning/repository/productRepository"
)

//сервис удаления товара
func (s *productService) DeleteProduct(product *productRepository.Product) {
	productRepository.Repository.Delete(product)
}
