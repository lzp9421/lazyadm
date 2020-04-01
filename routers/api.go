package routers

import (
	"github.com/lazygo/lazygo"
	"lazyadm/app/controller"
)

func RegisterApiHandler(router *lazygo.Router) {
	router.HandleFunc("/api/user/{action:[0-9a-zA-Z_]+}", &controller.User{})
	router.HandleFunc("/api/menu/{action:[0-9a-zA-Z_]+}", &controller.Menu{})
}
