package productService

import (
	"learning/processors/productProcessors/conditions"
)

type productService struct {
	processorsCond [5]processorsProductConditions.ProcessorProductConditionInterface
}


func newProductService() *productService  {
	var processorsCond [5]processorsProductConditions.ProcessorProductConditionInterface

	processorsCond[0] = processorsProductConditions.ProcessorCondWhere
	processorsCond[1] = processorsProductConditions.ProcessorCondSort
	processorsCond[2] = processorsProductConditions.ProcessorCondOrder
	processorsCond[3] = processorsProductConditions.ProcessorCondLimit
	processorsCond[4] = processorsProductConditions.ProcessorCondOffset

	return &productService{processorsCond}
}