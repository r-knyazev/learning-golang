package categoryService

import (
	"learning/processors/categoryProcessors/conditions"
	"learning/processors/categoryProcessors/validators"
)

type categoryService struct {
	processorsCond 		[]processorsCategoryConditions.ProcessorCategoryConditionInterface
	processorsValidate	[]processorsCategoryValidators.ProcessorCategoryValidateInterface
}


func newCategoryService() *categoryService  {
	processorsCond := []processorsCategoryConditions.ProcessorCategoryConditionInterface{
		processorsCategoryConditions.ProcessorCondSort,
		processorsCategoryConditions.ProcessorCondOrder,
		processorsCategoryConditions.ProcessorCondLimit,
		processorsCategoryConditions.ProcessorCondOffset,
	}

	processorsValid := []processorsCategoryValidators.ProcessorCategoryValidateInterface{
		processorsCategoryValidators.ProcessorValidateName,
	}

	return &categoryService{processorsCond, processorsValid}
}