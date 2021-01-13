package productController

import "github.com/valyala/fasthttp"

//интерфейс контроллера товара
type ProductControllerInterface interface {
	GetList(ctx *fasthttp.RequestCtx)
	CreateProduct(ctx *fasthttp.RequestCtx)
	UpdateProduct(ctx *fasthttp.RequestCtx)
	DeleteProduct(ctx *fasthttp.RequestCtx)
}