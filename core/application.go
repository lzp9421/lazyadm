package core

import (
	"errors"
	"github.com/revel/config"
	"lazyadm/core/database"
	"lazyadm/core/library"
)

type Application struct {
	di       map[string]interface{}
	conf     *config.Context
	mysql    *database.Mysql
	redis    *database.Redis
	memcache *database.Memcache
	router   *Router
	server   *Server
}

var application *Application

func App() *Application {
	if application == nil {
		application = &Application{}
	}
	return application
}

func (a *Application) Run() {
	server := a.GetServer()
	server.Listen()
}

func (a *Application) SetDi(name string, di interface{}) {
	a.di[name] = di
}

func (a *Application) GetDi(name string) interface{} {
	return a.di[name]
}

func (a *Application) InitConfig(filename string) bool {
	conf, err := config.LoadContext(filename, []string{""})
	library.CheckError(err)
	a.conf = conf
	return true
}

func (a *Application) InitMysql() bool {
	if a.conf == nil {
		panic(errors.New("配置信息未初始化"))
	}
	mysql, err := database.NewMysql(a.conf)
	library.CheckError(err)
	a.mysql = mysql
	return true
}

func (a *Application) GetMysql() *database.Mysql {
	if a.mysql == nil {
		panic(errors.New("数据库未初始化"))
	}
	return a.mysql
}

func (a *Application) InitRedis() bool {
	if a.conf == nil {
		panic(errors.New("配置信息未初始化"))
	}
	redis, err := database.NewRedis(a.conf)
	library.CheckError(err)
	a.redis = redis
	return true
}

func (a *Application) GetRedis() *database.Redis {
	if a.redis == nil {
		panic(errors.New("Redis未初始化"))
	}
	return a.redis
}

func (a *Application) InitMemcache() bool {
	if a.conf == nil {
		panic(errors.New("配置信息未初始化"))
	}
	memcache, err := database.NewMemcache(a.conf)
	library.CheckError(err)
	a.memcache = memcache
	return true
}

func (a *Application) GetMemcache() *database.Memcache {
	if a.memcache == nil {
		panic(errors.New("Memcache未初始化"))
	}
	return a.memcache
}

func (a *Application) InitRouter(route func(*Router)) bool {
	router, err := NewRouter()
	library.CheckError(err)
	a.router = router
	route(a.router)
	return true
}

func (a *Application) InitServer() bool {
	if a.conf == nil {
		panic(errors.New("配置信息未初始化"))
	}
	if a.router == nil {
		panic(errors.New("路由未初始化"))
	}
	server, err := NewServer(a.conf, a.router)
	library.CheckError(err)
	a.server = server
	return true
}

func (a *Application) GetServer() *Server {
	if a.server == nil {
		panic(errors.New("服务器未初始化"))
	}
	return a.server
}
