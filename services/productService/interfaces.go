package productService

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

type ProductServiceInterface interface {
	GetProduct(id uint) *productRepository.Product
	GetProducts(params *requestService.RequestParams) []productRepository.Product
	CreateProduct(params *requestService.RequestParams) (*productRepository.Product, error)
	UpdateProduct(params *requestService.RequestParams, product *productRepository.Product) error
}