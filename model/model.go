package model

import (
	"github.com/lazygo/lazygo/config"
	"github.com/lazygo/lazygo/mysql"
	"github.com/lazygo/lazygo/utils"
)

func Init() {
	mysqlConfig, err := config.GetSection("mysql")
	utils.CheckFatal(err)
	err = mysql.Init(mysqlConfig)
	utils.CheckFatal(err)
}

type Model struct {
}

func (m *Model) Db(dbName string) *mysql.Db {
	database, err := mysql.Database(dbName)
	utils.CheckFatal(err)
	return database
}
