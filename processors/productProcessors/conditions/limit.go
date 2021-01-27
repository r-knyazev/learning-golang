package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
	"strconv"
)

func (p processorCondLimit) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	if params.Limit <= 0 {
		limit, _ := strconv.Atoi(os.Getenv("DEFAULT_PRODUCT_LIMIT"))
		output.Limit = uint(limit)

		return
	}

	output.Limit = params.Limit
}