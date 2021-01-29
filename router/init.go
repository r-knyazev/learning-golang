package router

import (
	"github.com/fasthttp/router"
	"learning/controllers/categoryController"
	"learning/controllers/productController"
)

var Router *router.Router

func init() {
	Router = router.New()

	Router.PUT("/api/v1/product/", productController.Controller.UpdateProduct)
	Router.GET("/api/v1/product/", productController.Controller.GetList)
	Router.POST("/api/v1/product/", productController.Controller.CreateProduct)
	Router.DELETE("/api/v1/product/", productController.Controller.DeleteProduct)

	Router.PUT("/api/v1/category/", categoryController.Controller.UpdateCategory)
	Router.GET("/api/v1/category/", categoryController.Controller.GetList)
	Router.POST("/api/v1/category/", categoryController.Controller.CreateCategory)
	Router.DELETE("/api/v1/category/", categoryController.Controller.DeleteCategory)

}