package categoryController

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"learning/services/categoryService"
	"learning/services/requestService"
)

type controller struct {
	requestService requestService.RequestServiceInterface
	categoryService categoryService.CategoryServiceInterface
}


func newCategoryController() CategoryControllerInterface {
	return &controller{
		requestService : requestService.RequestService,
		categoryService : categoryService.CategoryService}
}

//получение категории если передан id, иначе список категорий
func (c *controller) GetList(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	if requestParams.ID > 0 {
		if category := c.categoryService.GetCategory(requestParams.ID); category != nil {
			ctx.Response.SetStatusCode(200)
			json.NewEncoder(ctx).Encode(category)

			return
		}

		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "category not found"})

		return
	}

	if categories := c.categoryService.GetCategories(requestParams); len(categories) > 0 {
		ctx.Response.SetStatusCode(200)
		json.NewEncoder(ctx).Encode(categories)

		return
	}

	ctx.Response.SetStatusCode(404)
	json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "categories not found"})
}

//создание категории
func (c *controller) CreateCategory(ctx *fasthttp.RequestCtx) {
	ctx.Response.Header.Add("Content-Type", "application/json")

	requestParams := c.requestService.GetRequestParams(ctx)
	category, err := c.categoryService.CreateCategory(requestParams)

	if err != nil {
		ctx.Response.SetStatusCode(500)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : err.Error()})

		return
	}

	ctx.Response.SetStatusCode(201)
	json.NewEncoder(ctx).Encode(category)
}

//обновление категории
func (c *controller) UpdateCategory(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	category := c.categoryService.GetCategory(requestParams.ID)
	if category == nil {
		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "category not found"})

		return
	}

	err := c.categoryService.UpdateCategory(requestParams, category)
	if err != nil {
		ctx.Response.SetStatusCode(500)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : err.Error()})

		return
	}

	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx).Encode(category)
}

//удаление категории
func (c *controller) DeleteCategory(ctx *fasthttp.RequestCtx) {
	requestParams := c.requestService.GetRequestParams(ctx)
	ctx.Response.Header.Add("Content-Type", "application/json")

	category := c.categoryService.GetCategory(requestParams.ID)
	if category == nil {
		ctx.Response.SetStatusCode(404)
		json.NewEncoder(ctx).Encode(map[string]interface{}{"error" : "category not found"})

		return
	}

	c.categoryService.DeleteCategory(category)

	ctx.Response.SetStatusCode(200)
	json.NewEncoder(ctx).Encode(category)

}