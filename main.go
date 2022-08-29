package main

import (
	"db2struct/cmd"
	"flag"
)

var dsn = flag.String("dsn", "", "mysql-dsn")
var table = flag.String("table", "", "table-name")
var output = flag.String("output", "./", "output-dir")
var orm = flag.String("o", "gorm", "orm:xorm|gorm")

func main() {
	flag.Parse()
	cmd.Receive(*dsn, *table, *output, *orm)
}
