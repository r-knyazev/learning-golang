package productController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"new/src/repository/productRepository"
)

type controller struct {

}


func newProductController() ProductControllerInterface {
	return &controller{}
}

//получение товара
func (c *controller) GetList(ctx *fasthttp.RequestCtx) {
	var code int
	response := make(map[string]interface{})

	id, err := ctx.QueryArgs().GetUint("id")
	if err == nil && id > 0 {
		product := productRepository.Repository.GetById(uint(id))
		if product != nil {
			response["product"] = product
			code = 200
		} else {
			response["error"] = "product not found"
			code = 404
		}
	} else {
		response["error"] = "error in incoming parameters"
		code = 500
	}

	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetStatusCode(code)

	json.NewEncoder(ctx).Encode(response)
}

func (c *controller) CreateProduct(ctx *fasthttp.RequestCtx) {

}

func (c *controller) UpdateProduct(ctx *fasthttp.RequestCtx) {

}

func (c *controller) DeleteProduct(ctx *fasthttp.RequestCtx) {

}
