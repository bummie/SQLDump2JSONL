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

		switch rows := statement.Rows.(type) {
		case sqlparser.Values:
			for _, row := range rows {
				for _, value := range row {
					switch rowValue := value.(type) {
					case *sqlparser.SQLVal:
						fmt.Println(rowValue.Type == sqlparser.IntVal)
						fmt.Println(rowValue.Val)
					}

				}

			}
		}
		fmt.Println(statement.Table.Name)

	default:
		fmt.Println(statement)
	}
}
