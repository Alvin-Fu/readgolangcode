package main

import (
	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"
	"log"
	"readruntime/practice/httptest/httpsever"
)

func main() {
	router := fasthttprouter.New()
	router.GET("/", httpsever.Index)
	router.GET("/hello", httpsever.Hello)
	router.GET("/get", httpsever.TestGet)
	router.GET("/post", httpsever.TestPost)
	log.Fatal(fasthttp.ListenAndServe(":8088", router.Handler))
}
