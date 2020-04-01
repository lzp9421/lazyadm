package controller

import (
	"github.com/lazygo/lazygo/library"
	"lazyadm/app/model"
)

type User struct {
	Controller
}

func (u *User) IndexAction() {
	mdlLog := model.NewLog()
	data := mdlLog.GetById(1)
	data["mc_get"] = mdlLog.HdMc().Get("aa")
	u.Response(0, "aaaaaaa", data)
	library.Error("ssss")

}

func (u *User) LoginAction() {

	data := map[string]interface{}{
		"token": "token",
		"user": map[string]interface{} {
			"username": "小明",
		},
	}
	u.Response(0, "aaaaaaa", data)
}
