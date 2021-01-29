package processorsCategoryConditions

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
)

func (p processorCondOffset) GetCond(params *requestService.RequestParams, output *categoryRepository.FindByConditions) {
	output.Offset = params.Offset
}