package processorsProductConditions

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
)

func (p processorCondWhere) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	where := map[string]interface{}{}

	if params.CategoryID > 0 {
		where["category_id"] = params.CategoryID
	}

	output.Where = where
}