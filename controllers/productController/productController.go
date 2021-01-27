package productController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"learning/repository/productRepository"
	"learning/services/productService"
	"learning/services/requestService"
	"os"
	"strconv"
	"strings"
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