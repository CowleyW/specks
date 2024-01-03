package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"

	_ "github.com/go-sql-driver/mysql"
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
	case RandomNumber:
		return r.Intn(1000), nil
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
