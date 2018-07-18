package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	userId int    `json:userId`
	role   string `json:role`
	name   string `json:name`
}

// DB設定 接続
func dbConnect() *gorm.DB {
	const DBMS = "mysql"
	const USER = "root"
	const PROTOCOL = "tcp(127.0.0.1:3306)"
	const DBNAME = "gosample"
	const CONNECT = USER + "@" + PROTOCOL + "/" + DBNAME
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return db
}
