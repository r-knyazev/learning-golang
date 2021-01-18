package requestService

import "github.com/valyala/fasthttp"

type RequestServiceInterface interface {
	GetRequestParams(ctx *fasthttp.RequestCtx) *RequestParams
}