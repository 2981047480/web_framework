package main

import "github.com/2981047480/web_framework/framework"

func registerRouter(core *framework.Core) {
	core.Get("foo", FooControllerHandler)
}
