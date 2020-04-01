package routers

import (
	"github.com/lazygo/lazygo"
	"lazyadm/app/controller"
)

func RegisterWebHandler(router *lazygo.Router) {
	router.HandleFunc("/", &controller.Index{})
}
