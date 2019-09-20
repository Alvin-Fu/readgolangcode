package main

import (
	"httptest/httpsever"
	"log"
	"net/http"
	"time"
)

func main() {
	//http.HandleFunc("/hello", httpsever.HelloServer)
	httpsever.HandlersMap["/hello"] = httpsever.HelloW
	httpsever.HandlersMap["/f1"] = httpsever.F1
	httpsever.HandlersMap["/f2"] = httpsever.F2

	svr := &http.Server{
		Addr:           ":8080",
		Handler:        &httpsever.WebSever{},
		WriteTimeout:   10 * time.Second,
		ReadTimeout:    10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err := svr.ListenAndServe()
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
