package router

import (
	"github.com/lazygo/lazygo"
	"lazyadm/controller/api"
)

func ApiRouter(router *lazygo.Router) {

	//绑定路由
	//http://127.0.0.1:8000/api/index/test1
	//http://{host}:{port}/{包名}/{控制器名}/{方法名}
	router.RegisterController(&api.UserController{})
	router.RegisterController(&api.MenuController{})

}

func init()  {
	registry["api"] = ApiRouter
}
