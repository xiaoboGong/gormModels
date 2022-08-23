package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Receive(dsn,table,output string) {
	connectDB(dsn)
	sd := new(ShowDesc)
	tableInfo := sd.getTableInfo(table)
	tableInfo.generate(output)
}

func connectDB(dsn string) {
	DB, _ = sql.Open("mysql", dsn)
	DB.SetConnMaxLifetime(100)
	DB.SetMaxIdleConns(10)
	if err := DB.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("数据库链接成功")
}
