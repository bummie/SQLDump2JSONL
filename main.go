package main

import (
	"bufio"
	"os"
)

func main() {

	testParsing()
}

func testParsing() {
	file, err := os.Open("testdata/small.sql")

	if err != nil {
		panic(err.Error())
	}

	fileReader := bufio.NewReader(file)
	for {
		b, err := ReadSqlStatements(fileReader, []byte(";\n"))
		if err != nil {
			break
		}

		ParseSql(string(b))
	}
}
