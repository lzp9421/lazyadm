package controller

import (
	"lazyadm/app/model"
)

type User struct {
	Controller
}

func (u *User)IndexAction()  {
	mdlLog := model.NewLog()
	data := mdlLog.GetById(1)
	data["mc_get"] = mdlLog.HdMc().Get("aa")
	u.Response(0, "aaaaaaa", data)
}

func (u *User)Index2Action()  {
	u.Display("index")
}
