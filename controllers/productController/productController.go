package productController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"new/src/repository/productRepository"
	"os"
	"strconv"
	"strings"
)

type controller struct {

}


func newProductController() ProductControllerInterface {
	return &controller{}
}

//получение товара если передан id, иначе список товаров
func (c *controller) GetList(ctx *fasthttp.RequestCtx) {
	var code int
	response := make(map[string]interface{})

	id, _ := ctx.QueryArgs().GetUint("id")
	if id > 0 {
		product := productRepository.Repository.GetById(uint(id))
		if product != nil {
			response["product"] = product
			code = 200
		} else {
			code = 404
			response["error"] = "product not found"
		}
	} else {
		limit, _ := ctx.QueryArgs().GetUint("limit")
		if limit < 0 {
			limit, _ = strconv.Atoi(os.Getenv("DEFAULT_LIMIT"))
		}

		offset, _ := ctx.QueryArgs().GetUint("offset")
		if offset < 0 {
			offset = 0
		}

		var sort string
		if ctx.QueryArgs().Has("sort") {
			sort = string(ctx.QueryArgs().Peek("sort"))
		} else {
			sort = os.Getenv("DEFAULT_SORT")
		}

		var order string
		if ctx.QueryArgs().Has("order") {
			order = string(ctx.QueryArgs().Peek("order"))
		} else {
			order = os.Getenv("DEFAULT_ORDER")
		}

		wheres := make(map[string]interface{})

		categoryId, _ := ctx.QueryArgs().GetUint("category_id")
		if categoryId > 0 {
			wheres["category_id"] = categoryId
		}

		products := productRepository.Repository.FindBy(productRepository.FindByConditions{
			Where	: wheres,
			Sort	: sort,
			Order	: order,
			Limit	: uint(limit),
			Offset	: uint(offset)})

		if len(products) > 0 {
			code = 200
			response["products"] = products
		} else {
			code = 404
			response["error"] = "products not found"
		}
		response["products"] = products
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
				code = 201
			} else {
				response["error"] = "error while creating product"
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
//			  4) Уникальность артикула
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

	limit, _ := strconv.Atoi(os.Getenv("DEFAULT_LIMIT"))
	products := productRepository.Repository.FindBy(productRepository.FindByConditions{
		Where	: map[string]interface{}{"sku" : product.SKU},
		Limit	: uint(limit),
		Sort	: os.Getenv("DEFAULT_SORT"),
		Order	: os.Getenv("DEFAULT_ORDER")})

	if len(products) > 1 || (len(products) == 1 && products[0].ID != product.ID) {
		errors["sku"] = "not unique"
	}

	if len(errors) == 0 {
		return nil
	} else {
		return errors
	}
}