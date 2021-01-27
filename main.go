package main

import (
	"github.com/valyala/fasthttp"
	_ "learning/config"
	"learning/router"
	"log"
	"os"
)

func main() {
	log.Fatal(fasthttp.ListenAndServe(":" + os.Getenv("SERVER_PORT"), router.Router.Handler))
}