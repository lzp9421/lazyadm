package routers

import (
	"lazyadm/app/controller"
	"lazyadm/core"
)

func RegisterWebHanler(router *core.Router) {
	router.HandleFunc("/",controller.NewUser)
	router.HandleFunc("/user/{action:[0-9a-zA-Z_]+}", controller.NewUser)
	router.HandleFunc("/{action:[0-9a-zA-Z_]+}", controller.NewUser)
}
