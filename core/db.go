package core

import (
	"database/sql"
	"fmt"
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

var MysqlLostConnection = []string{
	"server has gone away",
	"no connection to the server",
	"Lost connection",
	"is dead or not enabled",
	"Error while sending",
	"decryption failed or bad record mac",
	"server closed the connection unexpectedly",
	"SSL connection has been closed unexpectedly",
	"Error writing data to the connection",
	"Resource deadlock avoided",
	"Transaction() on null",
	"child connection forced to terminate due to client_idle_limit",
	"query_wait_timeout",
	"reset by peer",
	"Physical connection is not usable",
	"TCP Provider: Error code 0x68",
	"ORA-03114",
	"Packets out of order. Expected",
	"Adaptive Server connection failed",
	"Communication link failure",
	"connection refused",
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

// 获取查询构建器
func (d *Db)Query(query string) (*sql.Rows) {
	var err error
	var rows *sql.Rows
	for retry := 2; retry > 0; retry-- {
		rows, err = d.db.Query(query)
		if err == nil {
			return rows
		}
		if !library.ContainInArray(err.Error(), MysqlLostConnection) {
			break
		}
	}
	if rows != nil {
		rows.Close()
	}
	fmt.Println(fmt.Sprintf("%s %s", query, err.Error()))
	return nil
}


func (d *Db) GetAll(query string) []map[string]interface{} {
	rows := d.Query(query)
	defer rows.Close()
	outArr := ParseData(rows)

	return outArr
}

func (d *Db) GetOne(query string) map[string]interface{} {
	rows := d.Query(query)
	defer rows.Close()
	outArr := ParseRowData(rows)

	return outArr
}

// 解析结果集
func ParseData(rows *sql.Rows) []map[string]interface{} {
	var data []map[string]interface{} = nil

	columns, err := rows.Columns()
	library.CheckError(err)
	fCount := len(columns)
	fieldPtr := make([]interface{}, fCount)
	fieldArr := make([]sql.RawBytes, fCount)
	fieldToID := make(map[string]int64, fCount)
	for k, v := range columns {
		fieldPtr[k] = &fieldArr[k]
		fieldToID[v] = int64(k)
	}

	for rows.Next() {
		err = rows.Scan(fieldPtr...)
		library.CheckError(err)

		m := make(map[string]interface{}, fCount)

		for k, v := range fieldToID {
			if fieldArr[v] == nil {
				m[k] = ""
			} else {
				m[k] = string(fieldArr[v])
			}
		}
		data = append(data, m)
	}

	err = rows.Err()
	library.CheckError(err)
	return data
}

// 解析单行结果集
func ParseRowData(rows *sql.Rows) map[string]interface{} {
	if rows == nil {
		return nil
	}
	columns, err := rows.Columns()
	library.CheckError(err)
	fCount := len(columns)
	fieldPtr := make([]interface{}, fCount)
	fieldArr := make([]sql.RawBytes, fCount)
	fieldToID := make(map[string]int64, fCount)
	for k, v := range columns {
		fieldPtr[k] = &fieldArr[k]
		fieldToID[v] = int64(k)
	}

	m := make(map[string]interface{}, fCount)
	if rows.Next() {
		err = rows.Scan(fieldPtr...)
		library.CheckError(err)

		for k, v := range fieldToID {
			if fieldArr[v] == nil {
				m[k] = ""
			} else {
				m[k] = string(fieldArr[v])
			}
		}
	}
	err = rows.Err()
	library.CheckError(err)
	return m
}