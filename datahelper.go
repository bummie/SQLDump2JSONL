package main

import (
	"bytes"
	"strings"
)

type reader interface {
	ReadString(delim byte) (line string, err error)
}

func ReadSqlStatements(r reader, delim []byte) (line []byte, err error) {
	for {
		s := ""
		s, err = r.ReadString(delim[len(delim)-1])

		if strings.HasPrefix(strings.TrimLeft(s, " "), "--") {
			continue
		}

		line = append(line, []byte(s)...)

		if err != nil {
			return
		}

		if bytes.HasSuffix(line, delim) {
			return line, nil
		}
	}
}
