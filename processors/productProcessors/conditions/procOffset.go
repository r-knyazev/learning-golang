package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

func (p processorCondOffset) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	output.Offset = params.Offset
}