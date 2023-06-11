package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

type reader interface {
	ReadString(delim byte) (line string, err error)
}

func ReadSqlStatements(r reader, delimLinux []byte, delimWindows []byte) (line []byte, err error) {
	for {
		s := ""
		s, err = r.ReadString(delimWindows[len(delimWindows)-1])

		if strings.HasPrefix(strings.TrimLeft(s, " "), "--") {
			continue
		}

		line = append(line, []byte(s)...)

		if err != nil {
			if err != io.EOF {
				fmt.Fprintln(os.Stderr, err)
			}
			return
		}

		if bytes.HasSuffix(line, delimLinux) || bytes.HasSuffix(line, delimWindows) {
			return line, nil
		}
	}
}

func WriteRow(app Application, row string, filename string) {

	if len(app.OutputPath) == 0 {
		fmt.Fprintln(os.Stdout, row)
		return
	}

	outPath, err := filepath.Abs(app.OutputPath)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Did not understand the provided path:\n"+err.Error())
		os.Exit(2)
	}

	outPath = path.Join(outPath, filename)

	file, err := os.OpenFile(outPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Fprintln(os.Stderr, "Failed writing line: "+row+"\n"+err.Error())
		return
	}
	defer file.Close()

	if _, err := file.WriteString(row + "\n"); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
