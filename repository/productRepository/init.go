package productRepository

import "new/src/connection"

var Repository ProductRepositoryInterface

func init()  {
	Repository = newProductRepository()

	connection.Connection.GetDB().Debug().AutoMigrate(&Product{})
}