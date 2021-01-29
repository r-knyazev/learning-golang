package processorsCategoryValidators

import (
	"errors"
	"learning/services/requestService"
)

func (p processorValidateName) IsValid(params *requestService.RequestParams) bool {
	return params.Name != ""
}

func (p processorValidateName) GetError() error {
	return errors.New("name required")
}