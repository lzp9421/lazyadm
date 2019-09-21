package database

import (
	"errors"
	"github.com/gomodule/redigo/redis"
	"github.com/letsfire/redigo/mode"
	"github.com/letsfire/redigo/mode/alone"
	"github.com/letsfire/redigo/mode/cluster"
	"github.com/letsfire/redigo/mode/sentinel"
	"github.com/revel/config"
	"strings"
	"time"
)

type Redis struct {
	pool map[string]mode.IMode
	conf *config.Context
}

func NewRedis(conf *config.Context) (*Redis, error) {
	return &Redis{
		pool: map[string]mode.IMode{},
		conf: conf,
	}, nil
}

func (r *Redis) RedisPool(name string, db int) mode.IMode {
	if _, ok := r.pool[name]; !ok {
		conf := r.conf
		conf.SetSection("redis." + name)
		defer conf.SetSection("DEFAULT")

		modeConf := conf.StringDefault("mode", "alone")
		authConf := conf.StringDefault("auth", "")
		serverConf := conf.StringDefault("server", "127.0.0.1:11211")

		server := strings.Split(serverConf, ",")
		for index, val := range server {
			server[index] = strings.Trim(val, " ")
		}

		poolOpts := []mode.PoolOption{
			mode.MaxActive(0),       // 最大连接数，默认0无限制
			mode.MaxIdle(0),         // 最多保持空闲连接数，默认2*runtime.GOMAXPROCS(0)
			mode.Wait(false),        // 连接耗尽时是否等待，默认false
			mode.IdleTimeout(0),     // 空闲连接超时时间，默认0不超时
			mode.MaxConnLifetime(0), // 连接的生命周期，默认0不失效
			mode.TestOnBorrow(nil),  // 空间连接取出后检测是否健康，默认nil
		}
		dialOpts := []redis.DialOption{
			redis.DialReadTimeout(time.Second),    // 读取超时，默认time.Second
			redis.DialWriteTimeout(time.Second),   // 写入超时，默认time.Second
			redis.DialConnectTimeout(time.Second), // 连接超时，默认500*time.Millisecond
			redis.DialPassword(authConf),          // 鉴权密码，默认空
			redis.DialDatabase(db),                // 数据库号，默认0
			redis.DialKeepAlive(time.Minute * 5),  // 默认5*time.Minute
			redis.DialNetDial(nil),                // 自定义dial，默认nil
			redis.DialUseTLS(false),               // 是否用TLS，默认false
			redis.DialTLSSkipVerify(false),        // 服务器证书校验，默认false
			redis.DialTLSConfig(nil),              // 默认nil，详见tls.Config
		}

		if modeConf == "alone" {
			r.pool[name] = alone.New(
				alone.Addr(server[0]),
				alone.PoolOpts(poolOpts...),
				alone.DialOpts(dialOpts...),
			)
		} else if modeConf == "sentinel" {
			r.pool[name] = sentinel.New(
				sentinel.Addrs(server),
				// 这两项配置和Alone模式完全相同
				sentinel.PoolOpts(poolOpts...),
				sentinel.DialOpts(dialOpts...),
				// 连接哨兵配置，用法于sentinel.DialOpts()一致
				// 默认未配置的情况则直接使用sentinel.DialOpts()的配置
				// sentinel.SentinelDialOpts()
			)
		} else if modeConf == "cluster" {
			r.pool[name] = cluster.New(
				cluster.Nodes(server),
				// 这两项配置和Alone模式完全相同
				cluster.PoolOpts(poolOpts...),
				cluster.DialOpts(dialOpts...),
			)
		} else {
			panic(errors.New("redis mode not support"))
		}
	}
	return r.pool[name]
}
