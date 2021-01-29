package processorsProductValidators

var (
	ProcessorValidateCategoryId		ProcessorProductValidateInterface
	ProcessorValidateSKURequired	ProcessorProductValidateInterface
	ProcessorValidateSKUUnique		ProcessorProductValidateInterface
	ProcessorValidateName			ProcessorProductValidateInterface
	)

type (
	processorValidateCategoryId		struct {}
	processorValidateSKURequired	struct {}
	processorValidateSKUUnique		struct {}
	processorValidateName			struct {}
	)

func init()  {
	ProcessorValidateCategoryId		= processorValidateCategoryId{}
	ProcessorValidateSKURequired	= processorValidateSKURequired{}
	ProcessorValidateSKUUnique		= processorValidateSKUUnique{}
	ProcessorValidateName			= processorValidateName{}
}