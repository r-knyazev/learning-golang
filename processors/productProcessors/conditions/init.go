package processorsProductConditions

var (
	ProcessorCondWhere 	ProcessorProductConditionInterface
	ProcessorCondSort  	ProcessorProductConditionInterface
	ProcessorCondOrder 	ProcessorProductConditionInterface
	ProcessorCondLimit 	ProcessorProductConditionInterface
	ProcessorCondOffset	ProcessorProductConditionInterface
	)

type (
	processorCondWhere 	struct {}
	processorCondSort 	struct {}
	processorCondOrder 	struct {}
	processorCondLimit 	struct {}
	processorCondOffset	struct {}
	)

func init()  {
	ProcessorCondWhere 	= processorCondWhere{}
	ProcessorCondSort	= processorCondSort{}
	ProcessorCondOrder	= processorCondOrder{}
	ProcessorCondLimit	= processorCondLimit{}
	ProcessorCondOffset	= processorCondOffset{}
}