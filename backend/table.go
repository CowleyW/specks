package main

import (
	"errors"
)

type ColumnType struct {
	Name string `json:"name"`
}

type DataType interface {
	string | uint64 | float64
}

type Column struct {
	Name       string     `json:"columnName"`
	Type       ColumnType `json:"columnType"`
	PrimaryKey bool       `json:"columnPrimaryKey"`
	Unique     bool       `json:"columnUnique"`
}

type DataColumn[T DataType] struct {
	Name    string
	Type    ColumnType
	Entries []T
}

type Reference struct {
	Name        string `json:"referenceName"`
	TableIndex  uint16 `json:"tableIndex"`
	ColumnIndex uint16 `json:"columnIndex"`
	PrimaryKey  bool   `json:"referencePrimaryKey"`
	Unique      bool   `json:"referenceUnique"`
}

type Table struct {
	Name       string      `json:"tableName"`
	Columns    []Column    `json:"columns"`
	References []Reference `json:"references"`
}

type DataTable struct {
	Name        string
	DataColumns []interface{}
}

func GenerateTableData(table Table) (DataTable, error) {
	t := DataTable{
		Name:        table.Name,
		DataColumns: []interface{}{},
	}

	for _, col := range table.Columns {
		if c, err := generateColumnData(col); err != nil {
			return t, err
		} else {
			t.DataColumns = append(t.DataColumns, c)
		}
	}

	return t, nil
}

func generateColumnData(column Column) (interface{}, error) {
	switch column.Type.Name {
	case "Row Number":
		return generateRowNumberColumn(column), nil
	default:
		return nil, errors.New("unknown column type")
	}
}

func generateRowNumberColumn(column Column) DataColumn[uint64] {
	data := DataColumn[uint64]{
		Name:    column.Name,
		Type:    column.Type,
		Entries: []uint64{},
	}

	for i := uint64(0); i < 100; i += 1 {
		data.Entries = append(data.Entries, i)
	}

	return data
}
