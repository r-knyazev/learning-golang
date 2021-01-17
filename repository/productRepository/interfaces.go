package productRepository

type ProductRepositoryInterface interface {
	GetById(id uint) *Product
	FindBy(conditions FindByConditions) []Product
	Save(product *Product) (*Product, error)
}
