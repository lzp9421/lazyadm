package controller

import (
	"net/http"
	"lazyadm/app/model"
)

type User struct {
	Controller
}

func NewUser (res http.ResponseWriter, req *http.Request) (*User) {
	user := &User{}
	user.Init("user", user, res, req)
	return user
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
