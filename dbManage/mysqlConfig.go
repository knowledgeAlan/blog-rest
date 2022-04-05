package dbManage

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)


var DB *sqlx.DB


func InitDb() {

	db, err  := sqlx.Connect("mysql","article:root@(localhost:3306)/db?parseTime=true")
	if (err != nil){
		fmt.Println(err)
	}

	DB = db

 


}