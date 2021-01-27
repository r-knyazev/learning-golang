package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
)

func (p processorCondOrder) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	if params.Order == "" {
		output.Order = os.Getenv("DEFAULT_PRODUCT_ORDER")
		return
	}

	output.Order = params.Order
}