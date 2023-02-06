package dao

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var DB *sql.DB

func InitDB() {
	//访问本地数据库无密码//
	var db, err = sql.Open("mysql", "root:sglwh20041234@tcp(43.143.254.32:3306)/suning")
	if err != nil {
		log.Fatalf("connect mysql error : %v", err)
	}
	DB = db
	//检验连接//
	fmt.Println(db.Ping())
}
