package categoryService

import (
	"learning/repository/categoryRepository"
)

//сервис удаления категории
func (s *categoryService) DeleteCategory(category *categoryRepository.Category) {
	categoryRepository.Repository.Delete(category)
}
