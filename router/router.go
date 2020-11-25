package router

import "github.com/lazygo/lazygo"

type Register func (*lazygo.Router)

var registry = map[string]Register{}

func GetRouter() *lazygo.Router {
	var router = lazygo.NewRouter()

	for _, register := range registry {
		register(router)
	}

	return router
}