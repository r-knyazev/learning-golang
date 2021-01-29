package processorsCategoryConditions

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
	"os"
	"strconv"
)

func (p processorCondLimit) GetCond(params *requestService.RequestParams, output *categoryRepository.FindByConditions) {
	if params.Limit <= 0 {
		limit, _ := strconv.Atoi(os.Getenv("DEFAULT_PRODUCT_LIMIT"))
		output.Limit = uint(limit)

		return
	}

	output.Limit = params.Limit
}