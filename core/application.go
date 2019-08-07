package core

import (
	"fmt"
	"github.com/revel/config"
	"lazyadm/core/database"
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

func App() (*Application) {
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
	if err != nil {
		fmt.Println("配置文件读取失败")
		return false
	}
	a.conf = conf
	return true
}

func (a *Application) InitMysql() bool {
	if a.conf == nil {
		panic("配置信息未初始化")
		return false
	}
	mysql, err := database.NewMysql(a.conf)
	if err != nil {
		return false
	}
	a.mysql = mysql
	return true
}

func (a *Application) GetMysql() *database.Mysql {
	if a.mysql == nil {
		panic("数据库未初始化")
		return nil
	}
	return a.mysql
}

func (a *Application) InitRedis() bool {
	if a.conf == nil {
		panic("配置信息未初始化")
		return false
	}
	redis, err := database.NewRedis(a.conf)
	if err != nil {
		return false
	}
	a.redis = redis
	return true
}

func (a *Application) GetRedis() *database.Redis {
	if a.redis == nil {
		panic("Redis未初始化")
		return nil
	}
	return a.redis
}

func (a *Application) InitMemcache() bool {
	if a.conf == nil {
		panic("配置信息未初始化")
		return false
	}
	memcache, err := database.NewMemcache(a.conf)
	if err != nil {
		return false
	}
	a.memcache = memcache
	return true
}

func (a *Application) GetMemcache() *database.Memcache {
	if a.memcache == nil {
		panic("Memcache未初始化")
		return nil
	}
	return a.memcache
}

func (a *Application) InitRouter(route func(*Router)) bool {
	router, err := NewRouter()
	if err != nil {
		return false
	}
	a.router = router
	route(a.router)
	return true
}

func (a *Application) InitServer() bool {
	if a.conf == nil {
		panic("配置信息未初始化")
		return false
	}
	if a.router == nil {
		panic("路由未初始化")
		return false
	}
	server, err := NewServer(a.conf, a.router)
	if err != nil {
		panic("服务器初始化失败: " + err.Error())
		return false
	}
	a.server = server
	return true
}

func (a *Application) GetServer() *Server {
	if a.server == nil {
		panic("服务器未初始化")
		return nil
	}
	return a.server
}
