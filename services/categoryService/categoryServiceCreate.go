package categoryService

import (
	"errors"
	"learning/repository/categoryRepository"
	"learning/services/requestService"
	"log"
)

//сервис создрания категории
func (s *categoryService) CreateCategory(params *requestService.RequestParams) (*categoryRepository.Category, error) {
	category := &categoryRepository.Category{}

	for _, proc := range s.processorsValidate {
		if !proc.IsValid(params) {
			return category, proc.GetError()
		}
	}

	category.Name = params.Name
	category.Tags = params.Tags

	category, err := categoryRepository.Repository.Save(category)
	if err != nil {
		log.Println(err)

		return category, errors.New("error while save category")
	}

	return category, nil
}