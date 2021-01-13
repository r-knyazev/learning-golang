package productRepository

var Repository ProductRepositoryInterface

func init()  {
	Repository = newProductRepository()
}