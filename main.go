package main

import (
	"github.com/valyala/fasthttp"
	"log"
	_ "new/src/config"
	"new/src/router"
	"os"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router.Router.Handler))
}