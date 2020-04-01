package controller

type Index struct {
	Controller
}

func (i *Index) IndexAction() {
	i.Display("index")
}
