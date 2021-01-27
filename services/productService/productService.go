package productService

import (
	"learning/processors/productProcessors/conditions"
	processorsProductConditions2 "learning/processors/productProcessors/validators"
)

type productService struct {
	processorsCond 		[5]processorsProductConditions.ProcessorProductConditionInterface
	processorsValidate	[2]processorsProductConditions2.ProcessorProductValidateInterface
}


func newProductService() *productService  {
	var (
		processorsCond 	[5]processorsProductConditions.ProcessorProductConditionInterface
		processorsValid	[2]processorsProductConditions2.ProcessorProductValidateInterface
		)

	processorsCond[0] = processorsProductConditions.ProcessorCondWhere
	processorsCond[1] = processorsProductConditions.ProcessorCondSort
	processorsCond[2] = processorsProductConditions.ProcessorCondOrder
	processorsCond[3] = processorsProductConditions.ProcessorCondLimit
	processorsCond[4] = processorsProductConditions.ProcessorCondOffset

	processorsValid[0] = processorsProductConditions2.ProcessorValidateCategoryId
	processorsValid[1] = processorsProductConditions2.ProcessorValidateSKU

	return &productService{
		processorsCond,
		processorsValid}
}