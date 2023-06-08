package main

import (
	"bytes"
	"fmt"
	"os"
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
			return
		}

		if bytes.HasSuffix(line, delimLinux) || bytes.HasSuffix(line, delimWindows) {
			return line, nil
		}
	}
}

func WriteRowToFile(row string, filename string) {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)

	if err != nil {
		fmt.Println("Failed writing line: " + row)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(row + "\n"); err != nil {
		fmt.Println(err.Error())
	}
}
