package model

import "lazyadm/core"

type Model struct {
	core.Model
}

func (m *Model)HdDb() *core.Db {
	return m.Db("hd")
}

func (m *Model)HdDbRead() *core.Db {
	return m.Db("hd.read")
}