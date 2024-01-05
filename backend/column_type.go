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
	case Boolean:
		if r.Intn(1000) >= 500 {
			return true, nil
		} else {
			return false, nil
		}
	case FirstName:
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM first_names").Scan(&count)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed to scan query result for count")
		}

		random := r.Intn(count) + 1
		var name string
		err = db.QueryRow("SELECT name FROM first_names WHERE id = ?", random).Scan(&name)
		if err != nil {

			fmt.Println("Random number: ", random)
			return nil, errors.New("failed to scan query result for name")
		}

		return name, nil
	case LastName:
		var count int
		err := db.QueryRow("SELECT COUNT(*) FROM last_names").Scan(&count)
		if err != nil {
			fmt.Println(err)
			return nil, errors.New("failed to scan query result for count")
		}

		random := r.Intn(count) + 1
		var name string
		err = db.QueryRow("SELECT name FROM last_names WHERE id = ?", random).Scan(&name)
		if err != nil {
			fmt.Println("Random number: ", random)
			return nil, errors.New("failed to scan query result for name")
		}

		return name, nil
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
		return nil, errors.New("not implemented yet")
	case Boolean:
		return basic, nil
	case SSN:
		return nil, errors.New("not implemented yet")
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
