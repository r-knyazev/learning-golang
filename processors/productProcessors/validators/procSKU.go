package processorsProductConditions

import (
	"errors"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
	"strconv"
)

var skuError string

func (p processorValidateSKU) IsValid(params *requestService.RequestParams) bool {
	if params.SKU == "" {
		skuError = "sku required"

		return false
	}

	limit, _ := strconv.Atoi(os.Getenv("DEFAULT_PRODUCT_LIMIT"))
	products := productRepository.Repository.FindBy(productRepository.FindByConditions{
		Where	: map[string]interface{}{"sku" : params.SKU},
		Limit	: uint(limit),
		Sort	: os.Getenv("DEFAULT_PRODUCT_SORT"),
		Order	: os.Getenv("DEFAULT_PRODUCT_ORDER")})

	if len(products) > 1 || (len(products) == 1 && products[0].ID != params.ID) {
		skuError = "sku not unique"

		return false
	}

	return true
}

func (p processorValidateSKU) GetError() error {
	return errors.New(skuError)
}