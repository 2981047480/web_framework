package demo_test

import (
	"fmt"
	"testing"

	"github.com/2981047480/web_framework/demo"
	"github.com/gin-gonic/gin"
)

func TestDemo(t *testing.T) {
	demo.Demo()
}

func TestTest(t *testing.T) {
	var ctx *gin.Context
	c := ctx.Request.Context()
	// 这里的request就是http的request
	// type Context struct {
	//		...
	//		Request   *http.Request
	// 		...
	// }
	fmt.Println(c)
}
