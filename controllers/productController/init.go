package productController

var Controller ProductControllerInterface

//инициализация контроллера товара
func init() {
	Controller = newProductController()
}