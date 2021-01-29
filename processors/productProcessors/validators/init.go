package processorsProductValidators

var (
	ProcessorValidateCategoryId		ProcessorProductValidateInterface
	ProcessorValidateCategoryExist	ProcessorProductValidateInterface
	ProcessorValidateSKURequired	ProcessorProductValidateInterface
	ProcessorValidateSKUUnique		ProcessorProductValidateInterface
	ProcessorValidateName			ProcessorProductValidateInterface
	)

type (
	processorValidateCategoryId		struct {}
	processorValidateCategoryExist	struct {}
	processorValidateSKURequired	struct {}
	processorValidateSKUUnique		struct {}
	processorValidateName			struct {}
	)

func init()  {
	ProcessorValidateCategoryId		= processorValidateCategoryId{}
	ProcessorValidateCategoryExist	= processorValidateCategoryExist{}
	ProcessorValidateSKURequired	= processorValidateSKURequired{}
	ProcessorValidateSKUUnique		= processorValidateSKUUnique{}
	ProcessorValidateName			= processorValidateName{}
}