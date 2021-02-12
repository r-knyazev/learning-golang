package productService

import (
	"learning/repository/productRepository"
	"learning/services/requestService"
	"sync"
)

//сервис создрания товара
func (s *productService) CreateProduct(params *requestService.RequestParams) (*productRepository.Product, []string) {
	product := &productRepository.Product{}
	countProc := len(s.processorsValidate)
	itemsChan := make(chan error, countProc)

	var wg sync.WaitGroup

	defer close(itemsChan)

	wg.Add(2)
	for i := 1; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			for j := i; j < countProc; j = j + 2 {
				if !s.processorsValidate[j].IsValid(params) {
					itemsChan <- s.processorsValidate[j].GetError()
				}
			}
		}(i)
	}

	wg.Wait()

	countErrors := len(itemsChan)
	errors := make([]string, countErrors)

	for i := 0; i < countErrors; i++ {
		e := <- itemsChan
		errors[i] = e.Error()
	}

	if len(errors) > 0 {
		return nil, errors
	}

	product.Name = params.Name
	product.SKU = params.SKU
	product.CategoryID = params.CategoryID

	return product, nil
}