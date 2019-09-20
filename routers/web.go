package routers

import (
	"lazyadm/app/controller"
	"lazyadm/core"
)

func RegisterWebHanler(router *core.Router) {
	router.HandleFunc("/", &controller.User{})
	router.HandleFunc("/user/{action:[0-9a-zA-Z_]+}", &controller.User{})
	router.HandleFunc("/{action:[0-9a-zA-Z_]+}", &controller.User{})
}
