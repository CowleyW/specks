package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
)

type DataType string

const (
	FirstName    DataType = "First Name"
	LastName              = "Last Name"
	Character             = "Character"
	Age                   = "Age"
	Color                 = "Color"
	Boolean               = "Boolean"
	SSN                   = "SSN"
	RowNumber             = "Row Number"
	RandomNumber          = "Random Number"
	Date                  = "Date"
	Time                  = "Time"
	DateTime              = "DateTime"
)

type IColumnType interface {
	GenerateEntry(desc TableDesc, rowNumber uint, r *rand.Rand, db *sql.DB) (any, error)
}

type BasicColumnType struct {
	Name DataType `json:"name"`
}

func (bct BasicColumnType) GenerateEntry(desc TableDesc, rowNumber uint, r *rand.Rand, db *sql.DB) (any, error) {
	switch bct.Name {
	case RowNumber:
		return rowNumber, nil
	case Color:
		return queryRandomColumn(db, r, "colors", "name")
	case SSN:
		return fmt.Sprintf("%03d-%02d-%04d", r.Intn(1000), r.Intn(100), r.Intn(10000)), nil
	case Boolean:
		if r.Intn(1000) >= 500 {
			return true, nil
		} else {
			return false, nil
		}
	case FirstName:
		return queryRandomColumn(db, r, "first_names", "name")
	case LastName:
		return queryRandomColumn(db, r, "last_names", "name")
	default:
		return nil, errors.New("unknown column data type")
	}
}

type BoundedColumnType struct {
	Name DataType `json:"name"`
	Min  int      `json:"min"`
	Max  int      `json:"max"`
}

func (bct BoundedColumnType) GenerateEntry(desc TableDesc, rowNumber uint, r *rand.Rand, db *sql.DB) (any, error) {
	switch bct.Name {
	case Age:
		return r.Intn(bct.Max-bct.Min) + bct.Min, nil
	case RandomNumber:
		return r.Intn(bct.Max-bct.Min) + bct.Min, nil
	default:
		return nil, errors.New("unknown column data type")
	}
}

func ConstructColumnType(data json.RawMessage) (IColumnType, error) {
	var basic BasicColumnType
	if err := json.Unmarshal(data, &basic); err != nil {
		return nil, err
	}

	switch basic.Name {
	case FirstName:
		return basic, nil
	case LastName:
		return basic, nil
	case Character:
		return nil, errors.New("not implemented yet")
	case Age:
		return unmarshalAsBounded(data)
	case Color:
		return basic, nil
	case Boolean:
		return basic, nil
	case SSN:
		return basic, nil
	case RowNumber:
		return basic, nil
	case RandomNumber:
		return unmarshalAsBounded(data)
	case Date:
		return nil, errors.New("not implemented yet")
	case Time:
		return nil, errors.New("not implemented yet")
	case DateTime:
		return nil, errors.New("not implemented yet")
	default:
		return nil, errors.New("unknown column data type")
	}
}

func unmarshalAsBounded(data json.RawMessage) (IColumnType, error) {
	var bounded BoundedColumnType
	if err := json.Unmarshal(data, &bounded); err != nil {
		return nil, err
	} else {
		return bounded, nil
	}
}
