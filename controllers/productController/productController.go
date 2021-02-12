package productController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"learning/repository/productRepository"
	"learning/services/productService"
	"learning/services/requestService"
)

type controller struct {
	requestService requestService.RequestServiceInterface
	productService productService.ProductServiceInterface
}


func newProductController() ProductControllerInterface {
	return &controller{
		requestService : requestService.RequestService,
		productService : productService.ProductService}
}

//получение товара если передан id, иначе список товаров
func (c *controller) GetList(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	if requestParams.ID > 0 {
		if product := c.productService.GetProduct(requestParams.ID); product != nil {
			ctx.Response.SetStatusCode(200)
			json.NewEncoder(ctx).Encode(product)

			return
		}

		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "product not found"})

		return
	}

	if products := c.productService.GetProducts(requestParams); len(products) > 0 {
		ctx.Response.SetStatusCode(200)
		json.NewEncoder(ctx).Encode(products)

		return
	}

	ctx.Response.SetStatusCode(404)
	json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "products not found"})
}

//создание товара
func (c *controller) CreateProduct(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Content-Type", "application/json")

	requestParams := c.requestService.GetRequestParams(ctx)
	requestParams.ID = 0

	product, errors := c.productService.CreateProduct(requestParams)

	if errors != nil {
		ctx.Response.SetStatusCode(500)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"errors" : errors})

		return
	}

	product, err := productRepository.Repository.Save(product)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "error while save product"})

		return
	}

	ctx.Response.SetStatusCode(201)
	json.NewEncoder(ctx).Encode(product)
}

//обновление товара
func (c *controller) UpdateProduct(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	product := c.productService.GetProduct(requestParams.ID)
	if product == nil {
		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "product not found"})

		return
	}

	err := c.productService.UpdateProduct(requestParams, product)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : err.Error()})

		return
	}

	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx).Encode(product)
}

//удаление товара
func (c *controller) DeleteProduct(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	product := c.productService.GetProduct(requestParams.ID)
	if product == nil {
		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "product not found"})

		return
	}

	c.productService.DeleteProduct(product)

	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx).Encode(product)

}