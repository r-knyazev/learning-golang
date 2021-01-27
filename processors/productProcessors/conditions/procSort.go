package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
)

func (p processorCondSort) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	if params.Sort == "" {
		output.Sort = os.Getenv("DEFAULT_PRODUCT_SORT")
		return
	}

	output.Sort = params.Sort
}