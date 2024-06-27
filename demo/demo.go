package demo

import (
	"net/http"
)

func Demo() {
	fs := http.FileServer(http.Dir("/Users/zhaozephyr"))
	/*
		fs对象如下：
		type fileHandler struct {
		root FileSystem
		}

		其中FileSystem为一个实现了Open方法的接口
		type FileSystem interface {
		Open(name string) (File, error)
		}
	*/
	http.Handle("/static/", http.StripPrefix("/static", fs))
}
