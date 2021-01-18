package productService

import (
	"new/src/repository/productRepository"
	"new/src/services/requestService"
)

type ProductServiceInterface interface {
	GetProduct(id uint) *productRepository.Product
	GetProducts(params *requestService.RequestParams) []productRepository.Product
}