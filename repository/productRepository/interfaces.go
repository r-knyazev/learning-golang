package productRepository

type ProductRepositoryInterface interface {
	GetById(id uint) *Product
	Save(product *Product) (*Product, error)
}
