package productService

import (
	"learning/processors/productProcessors/conditions"
	processorsProductValidators "learning/processors/productProcessors/validators"
)

type productService struct {
	processorsCond 		[]processorsProductConditions.ProcessorProductConditionInterface
	processorsValidate	[]processorsProductValidators.ProcessorProductValidateInterface
}


func newProductService() *productService  {
	processorsCond := []processorsProductConditions.ProcessorProductConditionInterface{
		processorsProductConditions.ProcessorCondWhere,
		processorsProductConditions.ProcessorCondSort,
		processorsProductConditions.ProcessorCondOrder,
		processorsProductConditions.ProcessorCondLimit,
		processorsProductConditions.ProcessorCondOffset,
	}

	processorsValid := []processorsProductValidators.ProcessorProductValidateInterface{
		processorsProductValidators.ProcessorValidateCategoryId,
		processorsProductValidators.ProcessorValidateCategoryExist,
		processorsProductValidators.ProcessorValidateSKURequired,
		processorsProductValidators.ProcessorValidateSKUUnique,
		processorsProductValidators.ProcessorValidateName,
	}

	return &productService{processorsCond, processorsValid}
}