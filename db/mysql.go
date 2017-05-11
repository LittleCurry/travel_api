package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine

func InitMySQL(dsn string) {
	engine, _ = xorm.NewEngine("mysql", dsn)
	engine.DB().SetMaxOpenConns(100)
	engine.ShowSQL(true)
}

func MySQL() *xorm.Engine {
	return engine
}
