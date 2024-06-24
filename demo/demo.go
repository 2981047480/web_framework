package demo

import "net/http"

func Demo() {
	fs := http.FileServer(http.Dir("/Users/zhaozephyr"))
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
