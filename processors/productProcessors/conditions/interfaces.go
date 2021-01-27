package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

type ProcessorProductConditionInterface interface {
	GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions)
}