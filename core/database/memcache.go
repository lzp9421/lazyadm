package database

import (
	"github.com/revel/config"
	"github.com/bradfitz/gomemcache/memcache"
	"strings"
	"time"
)

type Memcache struct {
	mc map[string]*memcache.Client
	conf *config.Context
}

func NewMemcache(conf *config.Context) (*Memcache, error) {
	return &Memcache{
		mc:   map[string]*memcache.Client{},
		conf: conf,
	}, nil
}

func (m *Memcache) Mc(name string) *memcache.Client {
	if _, ok := m.mc[name]; ok {
		return m.mc[name]
	}

	conf := m.conf
	conf.SetSection("memcache." + name)
	defer conf.SetSection("DEFAULT")

	// 127.0.0.1:11211,127.0.0.1:11212,127.0.0.1:11213
	serverConf := conf.StringDefault("server", "127.0.0.1:11211")
	server := strings.Split(serverConf, ",")
	for index, val := range server {
		server[index] = strings.Trim(val, " ")
	}

	mc := memcache.New(server...)
	mc.MaxIdleConns = 10 // 最大保持10个空闲连接
	mc.Timeout = time.Duration(10) * time.Second

	m.mc[name] = mc
	return mc
}