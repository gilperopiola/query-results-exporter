package main

import (
	"database/sql"
	"log"

	"github.com/gilperopiola/query-results-exporter/config"
	_ "github.com/go-sql-driver/mysql"
)

type MyDatabase struct {
	*sql.DB
}

func (db *MyDatabase) Setup(cfg config.MyConfig) {
	var err error
	db.DB, err = sql.Open(
		cfg.DATABASE.TYPE, cfg.DATABASE.USERNAME+":"+cfg.DATABASE.PASSWORD+"@tcp("+cfg.DATABASE.HOSTNAME+":"+
			cfg.DATABASE.PORT+")/"+cfg.DATABASE.SCHEMA+"?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	err = db.DB.Ping()
	if err != nil {
		log.Fatalf("error pinging database: %v", err)
	}
}
