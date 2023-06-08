package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {

	sqlFilepath := flag.String("p", "", "Specify the path to your *.sql dump file")
	flag.Parse()

	if *sqlFilepath == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	parseSqlDump(*sqlFilepath)
}

func parseSqlDump(filePath string) {
	filePath, err := filepath.Abs(filePath)

	if err != nil {
		fmt.Println("Did not understand the provided path:\n" + err.Error())
		os.Exit(2)
	}
	file, err := os.Open(filePath)

	if err != nil {
		panic(err.Error())
	}

	fileReader := bufio.NewReader(file)
	for {
		rawSqlStatement, err := ReadSqlStatements(fileReader, []byte(";\n"), []byte(";\r\n"))
		if err != nil {
			break
		}

		ParseSql(string(rawSqlStatement))
	}
}
