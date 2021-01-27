package processorsProductConditions

var (
	ProcessorValidateCategoryId	ProcessorProductValidateInterface
	ProcessorValidateSKU		ProcessorProductValidateInterface
	)

type (
	processorValidateCategoryId	struct {}
	processorValidateSKU		struct {}
	)

func init()  {
	ProcessorValidateCategoryId	= processorValidateCategoryId{}
	ProcessorValidateSKU		= processorValidateSKU{}
}