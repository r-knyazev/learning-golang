package processorsProductValidators

import (
	"errors"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
	"strconv"
)

func (p processorValidateSKUUnique) IsValid(params *requestService.RequestParams) bool {
	limit, _ := strconv.Atoi(os.Getenv("DEFAULT_PRODUCT_LIMIT"))
	products := productRepository.Repository.FindBy(productRepository.FindByConditions{
		Where: map[string]interface{}{"sku": params.SKU},
		Limit: uint(limit),
		Sort:  os.Getenv("DEFAULT_PRODUCT_SORT"),
		Order: os.Getenv("DEFAULT_PRODUCT_ORDER")})

	//найденных товаров не должно быть больше 1;
	//если найден один торвар, и id наденного равен id из входящих параметров, тогда валидно
	return !(len(products) > 1 || (len(products) == 1 && products[0].ID != params.ID))
}

func (p processorValidateSKUUnique) GetError() error {
	return errors.New("sku not unique")
}