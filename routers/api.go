package routers

import (
	"lazyadm/app/controller"
	"lazyadm/core"
)

func RegisterApiHanler(router *core.Router) {
	router.HandleFunc("/api/user/{action:[0-9a-zA-Z_]+}", controller.NewUser)
}

