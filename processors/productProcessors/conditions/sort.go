package processorsProductConditions

import (
	"fmt"
	"learning/repository/productRepository"
	"learning/services/requestService"
	"os"
)

func (p processorCondSort) GetCond(params *requestService.RequestParams, output *productRepository.FindByConditions) {
	if params.Sort == "" {
		output.Sort = os.Getenv("DEFAULT_PRODUCT_SORT")
		fmt.Println("test1")
		return
	}

	output.Sort = params.Sort
}