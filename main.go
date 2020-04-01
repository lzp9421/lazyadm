package main

import (
	"github.com/lazygo/lazygo"
	"lazyadm/app"
	"lazyadm/routers"
)

func init() {

	regRouter := func(router *lazygo.Router) {
		router.HandleStatic(app.Asset)
		routers.RegisterWebHandler(router)
		routers.RegisterApiHandler(router)
	}

	lazygo.App().InitApp("lazyadm", regRouter, app.Asset)
}

func main() {
	lazygo.App().Run()
}
