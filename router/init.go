package router

import (
	"github.com/fasthttp/router"
	"new/src/controllers/productController"
)

var Router *router.Router

func init() {
	Router = router.New()

	Router.PUT("/api/v1/product/", productController.Controller.UpdateProduct)
	Router.GET("/api/v1/product/", productController.Controller.GetList)
	Router.POST("/api/v1/product/", productController.Controller.UpdateProduct)
	Router.DELETE("/api/v1/product/", productController.Controller.DeleteProduct)

	/*router.PUT("/api/v1/category/", controllers.CreateCategory)
	router.GET("/api/v1/category/", controllers.GetCategory)
	router.POST("/api/v1/category/", controllers.UpdateCategory)
	router.DELETE("/api/v1/category/", controllers.DeleteCategory)*/

}