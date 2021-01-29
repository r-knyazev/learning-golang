package processorsProductValidators

import (
	"errors"
	"learning/services/requestService"
)

var skuRequiredError string

func (p processorValidateSKURequired) IsValid(params *requestService.RequestParams) bool {
	return params.SKU != ""
}

func (p processorValidateSKURequired) GetError() error {
	return errors.New("sku required")
}