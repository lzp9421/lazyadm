package controller

type IndexController struct {
	Controller
}

func (ctl *IndexController) IndexAction() {
	ctl.Display("index")
}
