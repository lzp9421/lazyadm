package database

import (
	"github.com/revel/config"
	"github.com/gomodule/redigo/redis"
	"fmt"
	"time"
)

type Redis struct {
	pool map[string]*redis.Pool
	conf *config.Context
}

func NewRedis(conf *config.Context) (*Redis, error) {
	return &Redis{
		pool: map[string]*redis.Pool{},
		conf: conf,
	}, nil
}

func (r *Redis) RedisPool(name string, db int) *redis.Pool {

	if _, ok := r.pool[name]; ok {
		return r.pool[name]
	}

	conf := r.conf
	conf.SetSection("redis." + name)
	defer conf.SetSection("DEFAULT")

	address := fmt.Sprintf(
		"%s:%d",
		conf.StringDefault("host", "127.0.0.1"),
		conf.IntDefault("port", 6379),
	)
	auth := conf.StringDefault("auth", "")

	// 建立连接池
	pool := &redis.Pool{
		MaxIdle:     16,
		MaxActive:   0,
		IdleTimeout: 240 * time.Second,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			conn, err := redis.Dial(
				"tcp",
				address,
				redis.DialPassword(auth),
				redis.DialDatabase(db),
				redis.DialConnectTimeout(5 * time.Second),
				redis.DialReadTimeout(1 * time.Second),
				redis.DialWriteTimeout(1 * time.Second),
				redis.DialKeepAlive(1 * time.Second),
			)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
	}

	r.pool[name] = pool
	return pool
}