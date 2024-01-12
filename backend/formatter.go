package main

import (
	"fmt"
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

func FormatAsCSV(data TableData, desc TableDesc) OutputFile {
	var file OutputFile

	keys := make([]string, len(desc.Columns)+len(desc.References))
	i := 0
	for _, col := range desc.Columns {
		keys[i] = col.Name
		i += 1
	}
	for _, ref := range desc.References {
		keys[i] = ref.Name
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
		fmt.Println(data.Entries)
		fmt.Println(len(data.Entries))
		for _, value := range data.Entries {
			switch v := value.(type) {
			case string:
				builder.WriteString(v)
			case uint:
				builder.WriteString(strconv.Itoa(int(v)))
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
			default:
				fmt.Printf("unknown type %T.\n", v)
			}

			i += 1
			if i != end {
				builder.WriteByte(',')
			} else {
				builder.WriteByte('\n')
			}
		}
	}

	file.Name = fmt.Sprintf("%s.csv", data.Name)
	file.Data = builder.String()

	return file
}

func FormatAsSQL(data TableData, desc TableDesc) OutputFile {
	var file OutputFile

	// Get the keys
	keys := make([]string, len(desc.Columns)+len(desc.References))
	i := 0
	for _, col := range desc.Columns {
		keys[i] = col.Name
		i += 1
	}
	for _, ref := range desc.References {
		keys[i] = ref.Name
		i += 1
	}
	end := i

	// INSERT INTO tableName (columns...) VALUES (values...)
	var builder strings.Builder
	for i, key := range keys {
		builder.WriteByte('`')
		builder.WriteString(key)
		builder.WriteByte('`')
		if i < end-1 {
			builder.WriteByte(',')
		}
	}
	beginStatement := fmt.Sprintf("INSERT INTO `%s` (%s) VALUES (", data.Name, builder.String())

	builder.Reset()

	// Write the SQL insert statements
	for _, data := range data.RowData {
		i = 0
		builder.WriteString(beginStatement)
		for _, value := range data.Entries {
			switch v := value.(type) {
			case string:
				builder.WriteByte('\'')
				builder.WriteString(strings.Replace(v, "'", "''", -1))
				builder.WriteByte('\'')
			case uint:
				builder.WriteString(strconv.Itoa(int(v)))
			case int:
				builder.WriteString(strconv.Itoa(v))
			case bool:
				if v {
					builder.WriteString("TRUE")
				} else {
					builder.WriteString("FALSE")
				}
			case nil:
				builder.WriteString("NULL")
			}

			i += 1
			if i != end {
				builder.WriteByte(',')
			} else {
				builder.WriteString(");\n")
			}
		}
	}

	file.Name = fmt.Sprintf("%s.sql", data.Name)
	file.Data = builder.String()

	return file
}
