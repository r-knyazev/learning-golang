package processorsProductValidators

import "learning/services/requestService"

// Подставка для тестирования
type ProcessorValidatorsMock struct {
	IsValidResult bool
	GetErrorResult error
}

//тестирование на валидность
func (p ProcessorValidatorsMock) IsValid(params *requestService.RequestParams) bool {
	return p.IsValidResult
}

func (p ProcessorValidatorsMock) GetError() error {
	return p.GetErrorResult
}
