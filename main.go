package main

import (
	"net/http"

	"github.com/2981047480/web_framework/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{
		Handler: core,
		Addr:    "127.0.0.1:8888",
	}
	server.ListenAndServe()
}
