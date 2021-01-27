package productRepository

import "learning/connection"

var Repository ProductRepositoryInterface

func init()  {
	Repository = newProductRepository()

	connection.Connection.GetDB().Debug().AutoMigrate(&Product{})
}