package controller

import (
	"github.com/lazygo/lazygo"
	"lazyadm/asset"
)

type Controller struct {
	lazygo.Controller
}

// 控制器初始化时调用
// 可以控制器中，重写此方法
func (ctl *Controller) Init() {
	ctl.InitTpl("view/", ".html", asset.Asset)
	ctl.ErrorHandler(lazygo.ErrNotFound, func(err error) {
		ctl.Abort(404, err.Error())
	})
}