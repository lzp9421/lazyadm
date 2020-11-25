package router

import (
	"github.com/lazygo/lazygo"
	"lazyadm/controller"
)

func WebRouter(router *lazygo.Router) {

	//绑定路由
	router.RegisterController(&controller.IndexController{})

}

func init()  {
	registry["web"] = WebRouter
}
