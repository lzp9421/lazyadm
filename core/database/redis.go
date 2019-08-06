package database

import (
	"github.com/revel/config"
	"github.com/gomodule/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
	conf *config.Context
}
