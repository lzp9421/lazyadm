package core

type Model struct {
}

var db = map[string]*Db{}
var mc = map[string]*Memcache{}

func (m *Model) Db(name string) *Db {
	if _, ok := db[name]; !ok {
		db[name] = NewDb(name)
	}
	return db[name]
}

func (m *Model) Mc(name string) *Memcache {
	if _, ok := mc[name]; !ok {
		mc[name] = NewMemcache(name)
	}
	return mc[name]
}
