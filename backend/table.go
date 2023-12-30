package main

import (
	"errors"
)

type ColumnType struct {
	Name string `json:"name"`
}

type DataType interface {
	string | int | uint | float64
}

type Column struct {
	Name       string     `json:"columnName"`
	Type       ColumnType `json:"columnType"`
	PrimaryKey bool       `json:"columnPrimaryKey"`
	Unique     bool       `json:"columnUnique"`
}

type DataColumn[T DataType] struct {
	Name    string
	Entries []T
}

type Reference struct {
	Name        string `json:"referenceName"`
	TableIndex  uint   `json:"tableIndex"`
	ColumnIndex uint   `json:"columnIndex"`
	PrimaryKey  bool   `json:"referencePrimaryKey"`
	Unique      bool   `json:"referenceUnique"`
}

type Table struct {
	Name       string      `json:"tableName"`
	Columns    []Column    `json:"columns"`
	References []Reference `json:"references"`
	NumRows    uint        `json:"numRows"`
}

type DataTable struct {
	Name     string                   `json:"name"`
	DataRows []map[string]interface{} `json:"rows"`
}

func GenerateTableData(table Table) (DataTable, error) {
	t := DataTable{
		Name:     table.Name,
		DataRows: []map[string]interface{}{},
	}

	// Generate Data for each column
	var columns []interface{}
	for _, col := range table.Columns {
		if c, err := generateColumnData(col, table.NumRows); err != nil {
			return t, err
		} else {
			columns = append(columns, c)
		}
	}

	// Convert the columns into rows
	for i := uint(0); i < table.NumRows; i += 1 {
		row := map[string]interface{}{}

		for _, column := range columns {
			switch col := column.(type) {
			case DataColumn[string]:
				row[col.Name] = col.Entries[i]
			case DataColumn[int]:
				row[col.Name] = col.Entries[i]
			case DataColumn[uint]:
				row[col.Name] = col.Entries[i]
			case DataColumn[float64]:
				row[col.Name] = col.Entries[i]
			default:
				return t, errors.New("unknown column type")
			}
		}

		t.DataRows = append(t.DataRows, row)
	}
	return t, nil
}

func generateColumnData(column Column, length uint) (interface{}, error) {
	switch column.Type.Name {
	case "Row Number":
		return generateRowNumberColumn(column, length), nil
	default:
		return nil, errors.New("unknown column type")
	}
}

func generateRowNumberColumn(column Column, length uint) DataColumn[uint] {
	data := DataColumn[uint]{
		Name:    column.Name,
		Entries: []uint{},
	}

	for i := uint(0); i < length; i += 1 {
		data.Entries = append(data.Entries, i)
	}

	return data
}
