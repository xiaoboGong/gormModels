package cmd

import (
	"os"
	"strings"
)
import "text/template"

const(
    StructTpl = `package {{.PackageName}}

type {{.StructName}} struct{
{{- range $i, $v := .FormatColumn}}
    {{$v.Field}} {{$v.Type}} {{$v.Tag}}
{{- end}}
}
`
	StructTpl2 = `package {{.PackageName}}

type {{.StructName}} struct{
{{- range $i, $v := .FormatColumn}}
    {{$v.Field}} {{$v.Type}} {{$v.Tag}}
{{- end}}
}

func (u {{.StructName}}) TableName() string {
	return "{{.OriTableName}}"
}
`
)

func (ts *TableInfo)generate(output string) {
	output = strings.TrimRight(output, "/")
	packName:= output
	if (output == "."){
		packName,_ = os.Getwd()
		output = ""
	}
	dir := strings.Split(packName,"/")
	ts.setPackageName(dir[len(dir)-1])
	tplStruct := StructTpl
	if ts.Orm == "xorm" {
		tplStruct = StructTpl2
	}
	tpl, err := template.New("struct").Parse(tplStruct)
	if err != nil{
		panic(err)
	}
	makeDir(output)
	handle := makeFile(output+"/"+ts.FileName+".go")
	tpl.Execute(handle,ts)
}

func makeDir(output string) {
	if output != "" {
		err := os.MkdirAll(output, 0766)
		if err != nil {
			panic(err.Error())
		}
	}
}

func makeFile(file string) *os.File {
	handle, _ := os.Create(file)
	return handle
}
