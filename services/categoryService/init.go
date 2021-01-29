package categoryService

var CategoryService CategoryServiceInterface

func init()  {
	CategoryService = newCategoryService()
}