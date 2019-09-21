package database

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/revel/config"
	"lazyadm/core/library"
	"time"
)

type Mysql struct {
	databases map[string]*sql.DB
	conf      *config.Context
}

func NewMysql(conf *config.Context) (*Mysql, error) {
	return &Mysql{
		databases: map[string]*sql.DB{},
		conf:      conf,
	}, nil
}

func (m *Mysql) Database(name string) *sql.DB {

	if _, ok := m.databases[name]; !ok {

		conf := m.conf
		conf.SetSection("mysql." + name)
		defer conf.SetSection("DEFAULT")

		dataSourceName := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&timeout=10s",
			conf.StringDefault("user", "root"),
			conf.StringDefault("passwd", "root"),
			conf.StringDefault("host", "127.0.0.1"),
			conf.StringDefault("port", "3306"),
			conf.StringDefault("dbname", "test"),
			conf.StringDefault("charset", "utf8"),
		)
		databases, err := sql.Open("mysql", dataSourceName)
		library.CheckError(err)
		databases.SetMaxIdleConns(16)
		databases.SetMaxOpenConns(conf.IntDefault("maxOpenConns", 1000))
		databases.SetConnMaxLifetime(time.Duration(60) * time.Second)

		m.databases[name] = databases
	}

	return m.databases[name]
}
