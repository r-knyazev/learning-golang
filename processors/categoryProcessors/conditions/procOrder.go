package processorsCategoryConditions

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
	"os"
)

func (p processorCondOrder) GetCond(params *requestService.RequestParams, output *categoryRepository.FindByConditions) {
	if params.Order == "" {
		output.Order = os.Getenv("DEFAULT_CATEGORY_ORDER")
		return
	}

	output.Order = params.Order
}