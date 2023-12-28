package main

type ColumnType struct {
	Name string `json:"name"`
}

type Column struct {
	ColumnName string     `json:"columnName"`
	ColumnType ColumnType `json:"columnType"`
	PrimaryKey bool       `json:"columnPrimaryKey"`
	Unique     bool       `json:"columnUnique"`
}

type Reference struct {
	ReferenceName string `json:"referenceName"`
	TableIndex    uint16 `json:"tableIndex"`
	ColumnIndex   uint16 `json:"columnIndex"`
	PrimaryKey    bool   `json:"referencePrimaryKey"`
	Unique        bool   `json:"referenceUnique"`
}

type Table struct {
	TableName  string      `json:"tableName"`
	Columns    []Column    `json:"columns"`
	References []Reference `json:"references"`
}
