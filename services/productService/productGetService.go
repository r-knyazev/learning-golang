package productService

import (
	"new/src/repository/productRepository"
	"new/src/services/requestService"
	"os"
	"strconv"
)

//получение товара по ID
//если товар не найден, вернет nil
func (s *productService) GetProduct(id uint) *productRepository.Product {
	return productRepository.Repository.GetById(id)
}

//получение товаров по фильтрам
func (s *productService) GetProducts(params *requestService.RequestParams) []productRepository.Product {
	return productRepository.Repository.FindBy(productRepository.FindByConditions{
		Where	: s.getWheres(params),
		Sort	: s.getSort(params),
		Order	: s.getOrder(params),
		Limit	: s.getLimit(params),
		Offset	: params.Offset})
}

//формирование where для запроса на выборку
func (s *productService) getWheres(params *requestService.RequestParams) map[string]interface{} {
	where := map[string]interface{}{}

	if params.CategoryID > 0 {
		where["category_id"] = params.CategoryID
	}

	return where
}

//возвращает Sort из RequestParams если передан, иначе дефолтное значение из конфига
func (s *productService) getSort(params *requestService.RequestParams) string {
	if params.Sort == "" {
		return os.Getenv("DEFAULT_PRODUCT_SORT")
	}

	return params.Sort
}

//возвращает Order из RequestParams если передан, иначе дефолтное значение из конфига
func (s *productService) getOrder(params *requestService.RequestParams) string {
	if params.Order == "" {
		return os.Getenv("DEFAULT_PRODUCT_ORDER")
	}

	return params.Order
}

//возвращает Limit из RequestParams если передан и > 0, иначе дефолтное значение из конфига
func (s *productService) getLimit(params *requestService.RequestParams) uint {
	if params.Limit <= 0 {
		limit, _ := strconv.Atoi(os.Getenv("DEFAULT_PRODUCT_LIMIT"))
		return uint(limit)

	}

	return params.Limit
}