package controller

import "net/http"

type User struct {
	Controller
}

func NewUser (res http.ResponseWriter, req *http.Request) (*User) {
	user := &User{}
	user.Init("user", user, res, req)
	return user
}

func (u *User)IndexAction()  {
	u.Response(0, "aaaaaaa", map[string]interface{}{"ssss": "dfasf"})
}

func (u *User)Index2Action()  {
	u.Display("index")
}
