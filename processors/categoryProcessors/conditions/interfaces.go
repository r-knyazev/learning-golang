package processorsCategoryConditions

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
)

type ProcessorCategoryConditionInterface interface {
	GetCond(params *requestService.RequestParams, output *categoryRepository.FindByConditions)
}