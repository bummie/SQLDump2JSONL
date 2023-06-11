package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Application struct {
	InputPath  string
	OutputPath string
}

func main() {

	sqlFilepath := flag.String("p", "", "Specify the path to your *.sql dump file")
	outputPath := flag.String("o", "", "Specify the output folder you want the resulting files to go")

	flag.Parse()

	app := &Application{
		InputPath:  *sqlFilepath,
		OutputPath: *outputPath,
	}

	parseSqlDump(app)
}

var EOL_LF = []byte(";\n")
var EOL_CRLF = []byte(";\r\n")

func parseSqlDump(app *Application) {

	inputReader := bufio.NewReader(os.Stdin)

	if len(app.InputPath) != 0 {
		// Reading from inputpath
		inputReader = createFileReader(app.InputPath)
	}

	for {
		rawSqlStatement, err := ReadSqlStatements(inputReader, EOL_LF, EOL_CRLF)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}

		ParseSql(*app, string(rawSqlStatement))
	}
}

func createFileReader(inputPath string) *bufio.Reader {
	filePath, err := filepath.Abs(inputPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Did not understand the provided path:\n"+err.Error())
		os.Exit(2)
	}

	file, err := os.Open(filePath)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		panic(err.Error())
	}

	return bufio.NewReader(file)
}
