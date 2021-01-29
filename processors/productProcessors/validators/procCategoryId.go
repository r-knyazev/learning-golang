package processorsProductValidators

import (
	"errors"
	"learning/services/requestService"
)

func (p processorValidateCategoryId) IsValid(params *requestService.RequestParams) bool {
	return params.CategoryID != 0
}

func (p processorValidateCategoryId) GetError() error {
	return errors.New("category_id required")
}