package main

import (
	"strconv"
	"strings"
)

type Format string

type OutputFile struct {
	Name, Data string
}

const (
	CSV  Format = "CSV"
	JSON        = "JSON"
	SQL         = "SQL"
)

func FormatAsCSV(data TableData) OutputFile {
	var file OutputFile

	keys := make([]string, len(data.RowData[0].Entries))
	i := 0
	for k := range data.RowData[0].Entries {
		keys[i] = k
		i += 1
	}
	end := i

	// Write the CSV header
	var builder strings.Builder
	for _, key := range keys {
		builder.WriteString(key)
		builder.WriteByte(',')
	}
	builder.WriteByte('\n')

	// Write the CSV data
	for _, data := range data.RowData {
		i = 0
		for _, value := range data.Entries {
			switch v := value.(type) {
			case string:
				builder.WriteString(v)
			case int:
				builder.WriteString(strconv.Itoa(v))
			case bool:
				if v {
					builder.WriteString("True")
				} else {
					builder.WriteString("False")
				}
			case nil:
				builder.WriteString("NULL")
			}

			i += 1
			if i != end {
				builder.WriteByte(',')
			} else {
				builder.WriteByte('\n')
			}
		}
	}

	file.Name = data.Name
	file.Data = builder.String()

	return file
}
