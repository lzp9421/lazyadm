package model

import "lazyadm/core"

var table = "log"

type Log struct {
	Model
	hdDb *core.Db
}

func NewLog() *Log {
	log := &Log{}
	log.hdDb = log.HdDb()
	return log
}

func (l *Log)Insert(data map[string]interface{}) int64 {
	table := l.hdDb.Table(table)
	return table.Insert(data)
}

func (l *Log)GetById(id int) map[string]interface{} {

	cond := map[string]interface{}{
		"id": 1,
	}
	table := l.hdDb.Table(table)
	return table.WhereMap(cond).FetchRow("*", "", "", 0)
}