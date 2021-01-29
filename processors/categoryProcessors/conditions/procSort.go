package processorsCategoryConditions

import (
	"learning/repository/categoryRepository"
	"learning/services/requestService"
	"os"
)

func (p processorCondSort) GetCond(params *requestService.RequestParams, output *categoryRepository.FindByConditions) {
	if params.Sort == "" {
		output.Sort = os.Getenv("DEFAULT_CATEGORY_SORT")
		return
	}

	output.Sort = params.Sort
}