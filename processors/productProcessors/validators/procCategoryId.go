package processorsProductConditions

import (
	"errors"
	"learning/services/requestService"
)

var categoryIdError string

func (p processorValidateCategoryId) IsValid(params *requestService.RequestParams) bool {
	if params.CategoryID == 0 {
		categoryIdError = "category_id required"

		return false
	}

	//TODO сделать проверку на существование категории

	return true
}

func (p processorValidateCategoryId) GetError() error {
	return errors.New(categoryIdError)
}