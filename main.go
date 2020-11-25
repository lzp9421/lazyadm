package main

import (
	"github.com/lazygo/lazygo"
	"github.com/lazygo/lazygo/config"
	"github.com/lazygo/lazygo/utils"
	"lazyadm/router"
)

func init() {
	config.LoadFile("lazyadm")
}

func main() {
	conf, err := config.GetSection("server")
	utils.CheckFatal(err)
	serv, err := lazygo.NewServer(conf, router.GetRouter())
	utils.CheckFatal(err)
	serv.Listen()
}
