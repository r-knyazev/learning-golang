package processorsCategoryConditions

var (
	ProcessorCondSort  	ProcessorCategoryConditionInterface
	ProcessorCondOrder 	ProcessorCategoryConditionInterface
	ProcessorCondLimit 	ProcessorCategoryConditionInterface
	ProcessorCondOffset	ProcessorCategoryConditionInterface
	)

type (
	processorCondSort 	struct {}
	processorCondOrder 	struct {}
	processorCondLimit 	struct {}
	processorCondOffset	struct {}
	)

func init()  {
	ProcessorCondSort	= processorCondSort{}
	ProcessorCondOrder	= processorCondOrder{}
	ProcessorCondLimit	= processorCondLimit{}
	ProcessorCondOffset	= processorCondOffset{}
}