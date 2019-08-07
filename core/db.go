package core

import (
	"database/sql"
	"lazyadm/core/library"
)

type Db struct {
	name string
	db   *sql.DB
}

// 分页返回数据 - 表字段名定义
type ResultRow struct {
}

// 分页返回数据 - 返回结果定义
type ResultData struct {
	List      []map[string]interface{}
	Count     int64
	PerPage   int
	Page      int
	PageCount int
	Start     int
	Mark      int
}

func NewDb(name string) *Db {
	db := App().GetMysql().Database(name)
	return &Db{
		name: name,
		db:   db,
	}
}

// 获取查询构建器
func (d *Db)Table(table string) *Builder {
	return NewBuilder(d, table)
}

func (d *Db) GetAll(query string) []map[string]interface{} {
	rows, err := d.db.Query(query)
	library.CheckQueryError(err, query)
	defer rows.Close()
	outArr := ParseData(rows)

	return outArr
}

func (d *Db) GetOne(query string) map[string]interface{} {
	rows, err := d.db.Query(query)
	library.CheckQueryError(err, query)
	defer rows.Close()
	outArr := ParseRowData(rows)

	return outArr
}
