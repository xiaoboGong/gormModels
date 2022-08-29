package cmd

import (
	"fmt"
	"regexp"
	"strings"
)

type ShowDesc struct {
	Field string
	Type string
	Null string
	Key string
	Default interface{}
	Extra string
}

type ColumnFormat struct {
	Field string
	Type string
	Tag string
}

type TableInfo struct {
	Orm string
	OriTableName string
	FileName string
	StructName string
	PackageName string
	StructInfo TableDesc
	FormatColumn []*ColumnFormat
}

type TableDesc []*ShowDesc

var TypeMysqlDicMp = map[string]string{
	"smallint":            "int16",
	"smallint unsigned":   "uint16",
	"int":                 "int",
	"int unsigned":        "uint",
	"bigint":              "int64",
	"bigint unsigned":     "uint64",
	"varchar":             "string",
	"char":                "string",
	"date":                "datatypes.Date",
	"datetime":            "time.Time",
	"bit(1)":              "[]uint8",
	"tinyint":             "int8",
	"tinyint unsigned":    "uint8",
	"tinyint(1)":          "bool", // tinyint(1) 默认设置成bool
	"tinyint(1) unsigned": "bool", // tinyint(1) 默认设置成bool
	"json":                "datatypes.JSON",
	"text":                "string",
	"timestamp":           "time.Time",
	"double":              "float64",
	"double unsigned":     "float64",
	"mediumtext":          "string",
	"longtext":            "string",
	"float":               "float32",
	"float unsigned":      "float32",
	"tinytext":            "string",
	"enum":                "string",
	"time":                "time.Time",
	"tinyblob":            "[]byte",
	"blob":                "[]byte",
	"mediumblob":          "[]byte",
	"longblob":            "[]byte",
	"integer":             "int64",
}

var TypeMysqlMatchList = []struct {
	Key   string
	Value string
}{
	{`^(tinyint)[(]\d+[)] unsigned`, "uint8"},
	{`^(smallint)[(]\d+[)] unsigned`, "uint16"},
	{`^(int)[(]\d+[)] unsigned`, "uint32"},
	{`^(bigint)[(]\d+[)] unsigned`, "uint64"},
	{`^(float)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(double)[(]\d+,\d+[)] unsigned`, "float64"},
	{`^(tinyint)[(]\d+[)]`, "int8"},
	{`^(smallint)[(]\d+[)]`, "int16"},
	{`^(int)[(]\d+[)]`, "int"},
	{`^(bigint)[(]\d+[)]`, "int64"},
	{`^(char)[(]\d+[)]`, "string"},
	{`^(enum)[(](.)+[)]`, "string"},
	{`^(varchar)[(]\d+[)]`, "string"},
	{`^(varbinary)[(]\d+[)]`, "[]byte"},
	{`^(blob)[(]\d+[)]`, "[]byte"},
	{`^(binary)[(]\d+[)]`, "[]byte"},
	{`^(decimal)[(]\d+,\d+[)]`, "float64"},
	{`^(mediumint)[(]\d+[)]`, "string"},
	{`^(double)[(]\d+,\d+[)]`, "float64"},
	{`^(float)[(]\d+,\d+[)]`, "float64"},
	{`^(datetime)[(]\d+[)]`, "time.Time"},
	{`^(bit)[(]\d+[)]`, "[]uint8"},
	{`^(text)[(]\d+[)]`, "string"},
	{`^(integer)[(]\d+[)]`, "int"},
	{`^(timestamp)[(]\d+[)]`, "time.Time"},
	{`^(geometry)[(]\d+[)]`, "[]byte"},
}

func (sd *ShowDesc) getTableInfo(tableName,orm string) (ti *TableInfo) {
	rows, err := DB.Query("desc "+ tableName)
	if err != nil {
		panic(err)
	}
	ti = new(TableInfo)
	for rows.Next() {
		b := new (ShowDesc)
		if e := rows.Scan(&b.Field,&b.Type,&b.Null,&b.Key,&b.Default,&b.Extra); e != nil {
			panic(e.Error())
		}
		ti.setTable(b)
	}
	ti.setName(tableName)
	ti.formatColumn(orm)
	rows.Close()
	return
}

func(ti *TableInfo)setName(tableName string) {
	ti.OriTableName = tableName
	names := strings.Split(tableName, "_")
	tempName := ""
	for _,v := range names {
		tempName += strings.ToUpper(v[:1])+strings.ToLower(v[1:])
	}
	ti.FileName,ti.StructName = strings.ToLower(tempName[:1])+tempName[1:],tempName
 }

func(ti *TableInfo) setTable(sd *ShowDesc) {
	temp := sd
	ti.StructInfo = append(ti.StructInfo, temp)
}

func (ti *TableInfo) setPackageName(name string) {
	ti.PackageName = name
}

func (ti *TableInfo) formatColumn(orm string) {
	ti.Orm = orm
	for _,v := range ti.StructInfo {
		// 字段
		fields := strings.Split(v.Field, "_")
		formatField := ""
		for _,v2 := range fields {
			formatField += strings.ToUpper(v2[:1]) + strings.ToLower(v2[1:])
		}

		// 类型
		formatType := ""
		//精确匹配
		if vType, ok := TypeMysqlDicMp[v.Type]; ok {
			formatType = vType
		}else {
			for _, l := range TypeMysqlMatchList {
				if ok, _ := regexp.MatchString(l.Key, v.Type); ok {
					formatType = l.Value
				}
			}
		}
		formatTag := ""
		// tag 设置
		if orm == "gorm" {
			formatTag = "`gorm:\"%scolumn:%s;not null\"`"
			f1 := ""
			if v.Key == "PRI" {
				f1 = "primaryKey;"
			}
			formatTag = strings.Trim(fmt.Sprintf(formatTag,f1,v.Field), " ")
		}else {
			formatTag = "`xorm:\"%s %s\"`"
			f1 := ""
			f2 := "notnull"
			if v.Key == "PRI" {
				f1 = "pk"
				if v.Extra == "auto_increment" {
					f1 += " autoincr"
				}
				f2 = ""
			}else {
				f1 = strings.Split(v.Type, " ")[0]
			}

			formatTag = strings.TrimRight(fmt.Sprintf(formatTag,f1,f2), " ")
		}

		ti.FormatColumn = append(ti.FormatColumn, &ColumnFormat{
			formatField,
			      formatType,
				  formatTag,
		})
	}
}

