package categoryController

var Controller CategoryControllerInterface

//инициализация контроллера товара
func init() {
	Controller = newCategoryController()
}