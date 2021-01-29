package categoryController

import "github.com/valyala/fasthttp"

//интерфейс контроллера товара
type CategoryControllerInterface interface {
	GetList(ctx *fasthttp.RequestCtx)
	CreateCategory(ctx *fasthttp.RequestCtx)
	UpdateCategory(ctx *fasthttp.RequestCtx)
	DeleteCategory(ctx *fasthttp.RequestCtx)
}