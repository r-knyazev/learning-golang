package categoryService

import (
	"errors"
	"learning/repository/categoryRepository"
	"learning/services/requestService"
	"log"
)

//сервис обновления категории
func (s *categoryService) UpdateCategory(params *requestService.RequestParams, category *categoryRepository.Category) error {

	for _, proc := range s.processorsValidate {
		if !proc.IsValid(params) {
			return proc.GetError()
		}
	}

	category.Name = params.Name
	category.Tags = params.Tags

	category, err := categoryRepository.Repository.Save(category)

	if err != nil {
		log.Println(err)

		return errors.New("error while save category")
	}

	return nil
}