package processorsProductValidators

import (
	"errors"
	"learning/repository/categoryRepository"
	"learning/services/requestService"
)

func (p processorValidateCategoryExist) IsValid(params *requestService.RequestParams) bool {
	return categoryRepository.Repository.GetById(params.CategoryID) != nil
}

func (p processorValidateCategoryExist) GetError() error {
	return errors.New("category not exist")
}