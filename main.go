package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Application struct {
	InputPath string
}

func main() {

	sqlFilepath := flag.String("p", "", "Specify the path to your *.sql dump file")
	flag.Parse()

	app := &Application{
		InputPath: *sqlFilepath,
	}

	parseSqlDump(app)
}

func parseSqlDump(app *Application) {

	inputReader := bufio.NewReader(os.Stdin)

	if len(app.InputPath) != 0 {
		// Reading from inputpath
		inputReader = createFileReader(app.InputPath)
	}

	for {
		rawSqlStatement, err := ReadSqlStatements(inputReader, []byte(";\n"), []byte(";\r\n"))
		if err != nil {
			break
		}

		ParseSql(string(rawSqlStatement))
	}
}

func createFileReader(inputPath string) *bufio.Reader {
	filePath, err := filepath.Abs(inputPath)

	if err != nil {
		fmt.Println("Did not understand the provided path:\n" + err.Error())
		os.Exit(2)
	}

	file, err := os.Open(filePath)

	if err != nil {
		panic(err.Error())
	}

	return bufio.NewReader(file)
}
