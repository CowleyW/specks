package main

import (
	"errors"
	rand2 "math/rand"
	"time"
)

type ColumnType struct {
	Name string `json:"name"`
}

type DataType interface {
	string | int | uint | float64 | bool
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
	var nonKeyColumns []interface{}
	for _, col := range table.Columns {
		if c, err := generateColumnData(col, table.NumRows); err != nil {
			return t, err
		} else {
			nonKeyColumns = append(nonKeyColumns, c)
		}
	}

	// Convert the columns into rows
	for i := uint(0); i < table.NumRows; i += 1 {
		row := map[string]interface{}{}

		for _, column := range nonKeyColumns {
			switch col := column.(type) {
			case DataColumn[string]:
				row[col.Name] = col.Entries[i]
			case DataColumn[int]:
				row[col.Name] = col.Entries[i]
			case DataColumn[uint]:
				row[col.Name] = col.Entries[i]
			case DataColumn[float64]:
				row[col.Name] = col.Entries[i]
			case DataColumn[bool]:
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
		return generateRowNumberColumn(column, length)
	case "Random Number":
		return generateRandomNumberColumn(column, length, 0, 1000)
	case "Boolean":
		return generateBooleanColumn(column, length, 50)
	default:
		return nil, errors.New("unknown column type")
	}
}

func generateRowNumberColumn(column Column, length uint) (DataColumn[uint], error) {
	data := DataColumn[uint]{
		Name:    column.Name,
		Entries: []uint{},
	}

	for i := uint(0); i < length; i += 1 {
		data.Entries = append(data.Entries, i)
	}

	return data, nil
}

// Generates a column containing boolean values.
//
// trueSkew should be a number in the range [0, 100] representing the approximate percentage of entries that are true
func generateBooleanColumn(column Column, length uint, trueSkew int) (DataColumn[bool], error) {
	rand := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	data := DataColumn[bool]{
		Name:    column.Name,
		Entries: []bool{},
	}

	if column.Unique && length > 2 {
		return data, errors.New("unique column has length greater than domain")
	}

	for i := uint(0); i < length; i += 1 {
		var val bool
		if rand.Intn(100) < trueSkew {
			val = true
		} else {
			val = false
		}
		data.Entries = append(data.Entries, val)
	}

	return data, nil
}

func generateRandomNumberColumn(column Column, length uint, min int, max int) (DataColumn[int], error) {
	rand := rand2.New(rand2.NewSource(time.Now().UnixNano()))
	data := DataColumn[int]{
		Name:    column.Name,
		Entries: []int{},
	}

	if column.Unique && length > uint(max-min) {
		return data, errors.New("unique column has length greater than domain")
	}

	for i := uint(0); i < length; i += 1 {
		data.Entries = append(data.Entries, rand.Intn(max-min)+min)
	}
	
	return data, nil
}
