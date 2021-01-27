package processorsProductConditions

import (
	"learning/services/requestService"
)

type ProcessorProductValidateInterface interface {
	IsValid(params *requestService.RequestParams) bool
	GetError() error
}