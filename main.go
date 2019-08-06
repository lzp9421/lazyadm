package main

import (
	"lazyadm/core"
	"lazyadm/routers"
	"lazyadm/app"
)

func init () {

	application := core.App()
	application.InitConfig("lazyadm.ini")
	application.InitMysql()
	application.InitRouter(func(router *core.Router) {
		router.HandleStatic(app.Asset)
		routers.RegisterWebHanler(router)
		routers.RegisterApiHanler(router)
	})
	application.InitServer()
}

func main () {
	core.App().Run()
}
