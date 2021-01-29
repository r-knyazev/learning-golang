package categoryRepository

import "learning/connection"

var Repository CategoryRepositoryInterface

func init()  {
	Repository = newCategoryRepository()

	connection.Connection.GetDB().Debug().AutoMigrate(&Category{})
}