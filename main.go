package main

import (
	"net/http"

	"github.com/2981047480/web_framework/framework"
)

func main() {
	server := &http.Server{
		Handler: framework.NewCore(),
		Addr:    ":8080",
	}
	server.ListenAndServe()
}
