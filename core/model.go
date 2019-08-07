package core

type Model struct {
}

var db = map[string]*Db{}

func (m *Model) Db(name string) *Db {
	if _, ok := db[name]; !ok {
		db[name] = NewDb(name)
	}
	return db[name]
}