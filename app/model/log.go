package model

var table string = "log"

type Log struct {
	Model
}

func (this *Log)Insert(data map[string]interface{}) int64 {

	return this.HdDb().Insert(table, data)
}