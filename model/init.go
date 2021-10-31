package model

import (
	_ "github.com/go-sql-driver/mysql"
	"fmt"
	"database/sql"
	"personal-site/log"
	"personal-site/utils"
)

var Db *sql.DB

func init() {
	var err error
	dsn:=fmt.Sprintf("%v:%v@tcp(%v%v)/%v",utils.Config.MysqlConf.User,utils.Config.MysqlConf.Password,utils.Config.MysqlConf.Addr,utils.Config.MysqlConf.Port,utils.Config.MysqlConf.Database)

	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Error(err.Error())
	}
	err = Db.Ping()
	if err != nil {
		log.Error("mysql连接失败：",err.Error())
		return
	}
	log.Info("mysql连接成功")
}
