package requestService

import (
	"github.com/valyala/fasthttp"
)

type RequestParams struct {
	ID			uint
	ProductID	uint
	CategoryID	uint
	Sort		string
	Order		string
	Limit		uint
	Offset		uint
}

//сбор всех доступных парамертов
func (s *requestService) GetRequestParams(ctx *fasthttp.RequestCtx) *RequestParams {
	id, _ 			:= ctx.QueryArgs().GetUint("id")
	productId, _ 	:= ctx.QueryArgs().GetUint("product_id")
	categoryId, _ 	:= ctx.QueryArgs().GetUint("category_id")
	sort			:= ctx.QueryArgs().Peek("sort")
	order 			:= ctx.QueryArgs().Peek("order")
	limit, _		:= ctx.QueryArgs().GetUint("limit")
	offset, _		:= ctx.QueryArgs().GetUint("offset")

	return &RequestParams{
		ID 			: s.intToUintOrZero(id),
		ProductID 	: s.intToUintOrZero(productId),
		CategoryID 	: s.intToUintOrZero(categoryId),
		Sort		: string(sort),
		Order		: string(order),
		Limit		: s.intToUintOrZero(limit),
		Offset		: s.intToUintOrZero(offset)}
}

//конвертирует int в uint
//если входящий параметр < 0, вернет 0
func (s *requestService) intToUintOrZero(val int) uint {
	if val < 0 {
		return 0
	}

	return uint(val)
}