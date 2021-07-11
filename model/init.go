package model

import (
	"database/sql"
	"personal-site/log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=go password=go dbname=ps_db sslmode=disable")
	if err != nil {
		log.Error(err.Error())
	}
}