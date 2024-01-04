package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
)

type ColumnDesc struct {
	Name       string      `json:"columnName"`
	Type       IColumnType `json:"columnType"`
	PrimaryKey bool        `json:"columnPrimaryKey"`
	Unique     bool        `json:"columnUnique"`
}

func (cd *ColumnDesc) UnmarshalJSON(data []byte) error {
	var temp struct {
		Name       string          `json:"columnName"`
		Type       json.RawMessage `json:"columnType"`
		PrimaryKey bool            `json:"columnPrimaryKey"`
		Unique     bool            `json:"columnUnique"`
	}

	if err := json.Unmarshal(data, &temp); err != nil {
		fmt.Println(err)
		return err
	}

	columnType, err := ConstructColumnType(temp.Type)
	if err != nil {
		return err
	}

	cd.Name = temp.Name
	cd.Type = columnType
	cd.PrimaryKey = temp.PrimaryKey
	cd.Unique = temp.Unique

	return nil
}

type ReferenceDesc struct {
	Name string `json:"referenceName"`

	// TableIndex must be less than or equal to the index of the table ReferenceDesc exists in.
	TableIndex uint   `json:"tableIndex"`
	ColumnName string `json:"columnName"`

	PrimaryKey bool `json:"referencePrimaryKey"`
	Unique     bool `json:"referenceUnique"`
	Default    any  `json:"default"`
}

func (rd ReferenceDesc) GenerateEntry(precursors []TableData, current TableData, r *rand.Rand) (any, error) {
	var referenceTable TableData
	if int(rd.TableIndex) < len(precursors) { // This ReferenceDesc references a previous TableData
		referenceTable = precursors[rd.TableIndex]
	} else if int(rd.TableIndex) == len(precursors) { // This ReferenceDesc references the current TableData
		referenceTable = current
	} else { // This ReferenceDesc references a future TableData, which is an error
		return nil, errors.New("table references future table")
	}

	tableLen := len(referenceTable.RowData)
	if tableLen == 0 {
		return rd.Default, nil
	} else {
		return referenceTable.RowData[r.Intn(tableLen)].Entries[rd.ColumnName], nil
	}
}

type TableDesc struct {
	Name       string          `json:"tableName"`
	Columns    []ColumnDesc    `json:"columns"`
	References []ReferenceDesc `json:"references"`
	NumRows    uint            `json:"numRows"`
}

func (t TableDesc) numPrimaryKeys() uint {
	count := uint(0)

	for _, column := range t.Columns {
		if column.PrimaryKey {
			count += 1
		}
	}

	for _, reference := range t.References {
		if reference.PrimaryKey {
			count += 1
		}
	}

	return count
}
