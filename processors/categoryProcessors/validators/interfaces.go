package processorsCategoryValidators

import (
	"learning/services/requestService"
)

type ProcessorCategoryValidateInterface interface {
	IsValid(params *requestService.RequestParams) bool
	GetError() error
}