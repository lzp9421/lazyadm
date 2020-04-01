package model

import (
	"github.com/lazygo/lazygo"
	"github.com/lazygo/lazygo/memcache"
	"github.com/lazygo/lazygo/mysql"
)

type Model struct {
	lazygo.Model
}

func (m *Model) HdDb() *mysql.Db {
	return m.Db("hd")
}

func (m *Model) HdDbRead() *mysql.Db {
	return m.Db("hd_read")
}

func (m *Model) HdMc() *memcache.Memcache {
	return m.Mc("hd")
}
