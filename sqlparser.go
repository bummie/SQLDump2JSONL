package main

import (
	"encoding/json"
	"fmt"
	"strconv"
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
		parseInsert(statement)
	default:
		fmt.Println(statement)
	}
}

func parseInsert(statement *sqlparser.Insert) {
	switch rows := statement.Rows.(type) {
	case sqlparser.Values:
		for _, row := range rows {
			jsonLine, err := rowToJsonLine(row, statement.Columns)

			if err != nil {
				fmt.Println("Failed writing line:\n" + err.Error())
			}

			WriteRowToFile(jsonLine, statement.Table.Name.String()+".jsonl")
		}
	}
}

func rowToJsonLine(row sqlparser.ValTuple, columns sqlparser.Columns) (string, error) {

	jsonRow := make(map[string]interface{})

	for rowIndex, value := range row {

		if len(columns)-1 < rowIndex {
			break
		}

		switch rowValue := value.(type) {
		case *sqlparser.SQLVal:
			switch rowValue.Type {
			case sqlparser.StrVal:
				jsonRow[columns[rowIndex].CompliantName()] = string(rowValue.Val)
			case sqlparser.IntVal:
				convertedValue, err := strconv.Atoi(string(rowValue.Val))
				if err != nil {
					return "", err
				}
				jsonRow[columns[rowIndex].CompliantName()] = convertedValue
			case sqlparser.FloatVal:
				convertedValue, err := strconv.ParseFloat(string(rowValue.Val), 64)
				if err != nil {
					return "", err
				}
				jsonRow[columns[rowIndex].CompliantName()] = convertedValue
			}
		}
	}

	out, err := json.Marshal(jsonRow)

	if err != nil {
		return "", err
	}

	return string(out), nil
}
