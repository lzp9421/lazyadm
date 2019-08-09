package core

import (
	"fmt"
	"lazyadm/core/library"
	"strings"
	"strconv"
	"math"
)

type Builder struct {
	schema *Db
	table  string
	cond   []string
}

func NewBuilder(schema *Db, table string) *Builder {
	return &Builder{
		schema: schema,
		table:  table,
		cond:   []string{},
	}
}

// Map查询
func (b *Builder) WhereMap(cond map[string]interface{}) *Builder {
	for k, v := range cond {
		b.cond = append(b.cond, build(k, "=", v))
	}
	return b
}

// 子句查询
func (b *Builder) WhereClause(cond string) *Builder {
	cond = strings.Trim(cond, " ")
	if cond != "" {
		b.cond = append(b.cond, cond)
	}
	return b
}

// IN查询
func (b *Builder) WhereIn(k string, in []string) *Builder {
	for i, v := range in {
		in[i] = library.Addslashes(v)
	}
	cond := fmt.Sprintf("`%s` IN('%s')", k, strings.Join(in, "', '"))
	b.cond = append(b.cond, cond)
	return b
}

// 构建条件
func (b *Builder) buildCond() string {
	if len(b.cond) == 0 {
		return ""
	}
	result := " WHERE " + strings.Join(b.cond, "AND")
	b.Clear()
	return result
}

// 清空当前where
func (b *Builder) Clear() *Builder {
	b.cond = []string{}
	return b
}

// 构建值
func BuildVal(val map[string]interface{}, extra []string) string {
	var items []string
	for k, v := range val {
		items = append(items, build(k, "=", v))
	}

	for _, v := range extra {
		v = strings.Trim(v, " ")
		if v != "" {
			items = append(items, v)
		}
	}
	result := strings.Join(items, ", ")
	return result
}

// 拼接查询语句字符串
func (b *Builder) MakeQueryString(fields string, order string, group string, limit int, start int) string {

	if b.table == "" {
		panic("没有指定表名")
	}

	queryString := fmt.Sprintf("SELECT %s FROM %s ", fields, b.table)

	cond := b.buildCond()
	if cond != "" {
		queryString += " " + cond
	}

	if group != "" {
		queryString += " GROUP BY " + group
	}

	if order != "" {
		queryString += " ORDER BY " + order
	}

	if limit > 0 && start > 0 {
		queryString += " LIMIT " + strconv.Itoa(start) + "," + strconv.Itoa(limit)
	} else if limit > 0 {
		queryString += " LIMIT " + strconv.Itoa(limit)
	}

	return queryString
}

//
func build(k string, op string, v interface{}) string {

	str := library.Addslashes(library.ToString(v))
	return fmt.Sprintf("`%s` %s '%s'", k, op, str)
}

// 构造分页返回结果
func Multi(count int64, perpage int, page int) ResultData {
	var data ResultData
	data.Count = int64(math.Max(float64(count), 0))
	data.Page = page
	data.PerPage = int(math.Max(float64(perpage), 1))
	data.PageCount = int(math.Ceil(float64(data.Count) / float64(data.PerPage)))
	data.Start = int(math.Max(float64(data.PerPage*data.Page-data.PerPage), 0))
	data.Mark = data.Start + 1
	return data
}

// 查询统计个数
func (b *Builder) Count() int64 {
	if b.table == "" {
		panic("没有指定表名")
	}
	where := b.buildCond()
	queryString := fmt.Sprintf("SELECT COUNT(*) AS num FROM %s %s", b.table, where)
	data := b.schema.GetOne(queryString)
	return library.ToInt64(data["num"])
}

// 查询并返回多条记录，且包含分页信息
func (b *Builder) FetchWithPage(fields string, order string, group string, limit int, page int) *ResultData {

	cond := b.buildCond()
	if len(cond) <= 0 {
		return &ResultData{}
	}
	count := b.Clear().WhereClause(cond).Count()
	data := Multi(count, limit, page)

	start := int(math.Max(float64(page * limit - limit), 0))

	data.List = b.Clear().WhereClause(cond).Fetch(fields, order, group, limit, start)
	return &data
}

// 查询并返回多条记录
func (b *Builder) Fetch(fields string, order string, group string, limit int, start int) []map[string]interface{} {

	queryString := b.MakeQueryString(fields, order, group, limit, start)

	rows := b.schema.Query(queryString)
	defer rows.Close()

	lists := ParseData(rows)

	return lists
}

// 查询并返回单条记录
func (b *Builder) FetchRow(fields string, order string, group string, start int) map[string]interface{} {

	queryString := b.MakeQueryString(fields, order, group, 1, start)

	rows := b.schema.Query(queryString)
	defer rows.Close()

	item := ParseRowData(rows)

	return item
}

// 查询并返回单个字段
func (b *Builder) FetchOne(field string, order string, group string, start int) string {

	queryString := b.MakeQueryString(field, order, group, 1, start)

	rows := b.schema.Query(queryString)
	defer rows.Close()

	item := ParseRowData(rows)

	return library.ToString(item[field])
}

// 插入
func (b *Builder) Insert(set map[string]interface{}) int64 {
	if b.table == "" {
		panic("没有指定表名")
	}

	queryString := "INSERT INTO " + b.table + " SET " + BuildVal(set, []string{})
	stmt, err := b.schema.db.Prepare(queryString)
	library.CheckError(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	library.CheckError(err)

	id, err := res.LastInsertId()
	library.CheckError(err)
	return id
}

// 更新
func (b *Builder) Update(set map[string]interface{}) int64 {
	if b.table == "" {
		panic("没有指定表名")
	}

	where := b.buildCond()
	if len(where) <= 0 {
		return 0
	}

	queryString := "UPDATE " + b.table + " SET " + BuildVal(set, []string{}) + where
	stmt, err := b.schema.db.Prepare(queryString)
	library.CheckError(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	library.CheckError(err)

	id, err := res.RowsAffected()
	library.CheckError(err)
	return id
}

// 自增
func (b *Builder) Increment(column string, amount int, set map[string]interface{}) int64 {
	if b.table == "" {
		panic("没有指定表名")
	}

	where := b.buildCond()
	if len(where) <= 0 {
		return 0
	}

	var extra []string
	if amount >= 0 {
		extra = []string{
			fmt.Sprintf("%s=%s+%d", column, column, amount),
		}
	} else {
		extra = []string{
			fmt.Sprintf("%s=%s-%d", column, column, -amount),
		}
	}

	queryString := "UPDATE " + b.table + " SET " + BuildVal(set, extra) + where
	stmt, err := b.schema.db.Prepare(queryString)
	library.CheckError(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	library.CheckError(err)

	id, err := res.RowsAffected()
	library.CheckError(err)
	return id
}

// 自减
func (b *Builder) Decrement(column string, amount int, set map[string]interface{}) int64 {
	return b.Increment(column, -amount, set)
}

// 删除
func (b *Builder) Delete() int64 {
	if b.table == "" {
		panic("没有指定表名")
	}
	where := b.buildCond()
	if len(where) <= 0 {
		return 0
	}

	queryString := "DELETE FROM " + b.table + " " + where
	stmt, err := b.schema.db.Prepare(queryString)
	library.CheckError(err)
	defer stmt.Close()

	res, err := stmt.Exec()
	library.CheckError(err)

	id, err := res.RowsAffected()
	library.CheckError(err)
	return id
}
