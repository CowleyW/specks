package main

type Column struct {
	name     string
	datatype DataType
}

type Table struct {
	name    string
	columns []Column
}
