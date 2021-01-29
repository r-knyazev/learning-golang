package categoryRepository

type CategoryRepositoryInterface interface {
	GetById(id uint) *Category
	FindBy(conditions FindByConditions) []Category
	Save(category *Category) (*Category, error)
	Delete(category *Category)
}
