package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/2981047480/web_framework/framework"
)

func FooControllerHandler(c *framework.Context) error {
	finish := make(chan struct{}, 1)
	panicChan := make(chan interface{}, 1)

	durationCtx, cancel := context.WithTimeout(c.BaseContext(), time.Duration(1*time.Second))
	defer cancel()
	// 这里读到了取消context的cancel方法

	go func() {
		defer func() {
			if p := recover(); p != nil {
				panicChan <- p
			}
		}()
		time.Sleep(2 * time.Second)
		c.JSON(200, "ok")

		finish <- struct{}{}
	}()

	select {
	// panic了
	case p := <-panicChan:
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		log.Println(p)
		c.JSON(500, "panic")
	// 正常
	case <-finish:
		fmt.Println("finish")
	// 如果子context返回done，则代表超时了
	case <-durationCtx.Done():
		c.WriterMux().Lock()
		defer c.WriterMux().Unlock()
		c.JSON(500, "timeout")
		c.SetHasTimeout()
	}
	return nil
}
