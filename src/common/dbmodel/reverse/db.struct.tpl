    // 数据库表结构定义
    package dbmodel

    {{$ilen := len .Imports}}
    {{if gt $ilen 0}}
    import (
      {{range .Imports}}"{{.}}"{{end}}
    )
    {{end}}

    {{range .Tables}}
    type {{TableMapper .Name}} struct {
    {{$table := .}}
    {{range .ColumnsSeq}} {{$col := $table.GetColumn .}} {{$colname := ColumnMapper $col.Name}} {{if eq $colname "CreatedAt"}} {{else if eq $colname "UpdatedAt"}} {{else if eq $colname "DeletedAt"}} {{else}} {{ColumnMapper $col.Name}} 	{{Type $col}} `{{Tag $table $col}}`
    {{end}} {{end}} CreatedAt time.Time `xorm:"created"`
    UpdatedAt time.Time `xorm:"updated"`
    DeletedAt time.Time `xorm:"deleted"`
    }
    {{end}}