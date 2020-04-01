package model

import (
	"github.com/lazygo/lazygo/library"
	"github.com/lazygo/lazygo/mysql"
)

var table = "log"

type Log struct {
	Model
	hdDb *mysql.Db
}

func NewLog() *Log {
	log := &Log{}
	log.hdDb = log.HdDb()
	return log
}

func (l *Log) Insert(data map[string]interface{}) int64 {
	table := l.hdDb.Table(table)
	res, err := table.Insert(data)
	library.CheckError(err)
	return res
}

func (l *Log) GetById(id int) map[string]interface{} {

	cond := map[string]interface{}{
		"id": 1,
	}
	table := l.hdDb.Table(table)
	res, err := table.WhereMap(cond).FetchRow("*", "", "", 0)
	library.CheckError(err)
	return res
}
