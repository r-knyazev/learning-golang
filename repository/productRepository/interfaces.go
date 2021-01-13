package productRepository

type ProductRepositoryInterface interface {
	GetById(id uint) *Product
}
