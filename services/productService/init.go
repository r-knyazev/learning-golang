package productService

var ProductService ProductServiceInterface

func init()  {
	ProductService = newProductService()
}