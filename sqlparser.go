package main

import (
	"fmt"
	"strings"

	"github.com/blastrain/vitess-sqlparser/sqlparser"
)

func ParseSql(rawSql string) {
	rawSql = strings.TrimSuffix(rawSql, "\n")
	statement, err := sqlparser.Parse(rawSql)
	if err != nil {
		fmt.Println(err)
		return
	}

	switch statement := statement.(type) {
	case *sqlparser.CreateTable:
		fmt.Println(statement.NewName.Name.String())
	case *sqlparser.Insert:
		fmt.Println("INSERT")
		fmt.Println(statement.Columns)
	default:
		fmt.Println(statement)
	}
}
