package productController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"new/src/repository/productRepository"
	"strings"
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

//создание товара
func (c *controller) CreateProduct(ctx *fasthttp.RequestCtx) {
	var code int
	product := &productRepository.Product{}
	response := make(map[string]interface{})

	r := strings.NewReader(string(ctx.Request.Body()))
	err := json.NewDecoder(r).Decode(product)

	if err == nil {
		validateErrors := c.validate(*product)
		if validateErrors == nil {
			createdProduct, err := productRepository.Repository.Save(product)
			if err == nil {
				response["created"] = createdProduct
				response["err"] = err
				code = 201
			} else {
				//TODO сформировать другой ответ в случае ошибки
				response["error"] = err
			}
		} else {
			response["error"] = validateErrors
			code = 500
		}
	} else {
		response["error"] = "error while decoding request body"
		response["detail"] = err
		response["msg"] = err.Error()
		code = 500
	}

	ctx.Response.Header.Add("Content-Type", "application/json")
	ctx.Response.SetStatusCode(code)

	json.NewEncoder(ctx).Encode(response)
}

func (c *controller) UpdateProduct(ctx *fasthttp.RequestCtx) {

}

func (c *controller) DeleteProduct(ctx *fasthttp.RequestCtx) {

}

// валидатор товаров
// проверяет: 1) Наличие category_id (обязательный параметр)
//            2) Существование категории TODO реализовать
//            3) Наличие артикула
// в случае успеха, возвращает nil
func (c *controller) validate(product productRepository.Product) map[string]interface{} {
	errors := make(map[string]interface{})

	if product.CategoryID == 0 {
		errors["category_id"] = "category_id required"
	}

	//TODO сделать проверку на существование категории

	if product.SKU == "" {
		errors["sku"] = "sku required"
	}

	//TODO сделать проверку на уникальность артикула

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}