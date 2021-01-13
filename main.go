package main

import (
	"github.com/valyala/fasthttp"
	"new/src/router"
)

func main() {
	//TODO вынести порт в конфиг
	fasthttp.ListenAndServe(":8080", router.Router.Handler)
}