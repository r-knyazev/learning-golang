package productRepository

type ProductRepositoryInterface interface {
	GetById(id uint) *Product
	FindBy(map[string]interface{}) []Product
	Save(product *Product) (*Product, error)
}
