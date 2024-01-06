package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"time"
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

type DateColumnType struct {
	Name   DataType   `json:"name"`
	Min    time.Time  `json:"min"`
	Max    time.Time  `json:"max"`
	Format DateFormat `json:"format"`
}

func (dct DateColumnType) GenerateEntry(desc TableDesc, rowNumber uint, r *rand.Rand, db *sql.DB) (any, error) {
	min := daysSinceEpoch(dct.Min)
	max := daysSinceEpoch(dct.Max)
	random := r.Intn(max-min) + min
	date := time.Unix(int64(random*86400), 0)

	switch dct.Format {
	case YYYY_MM_DD:
		return fmt.Sprintf("%04d-%02d-%02d", date.Year(), date.Month(), date.Day()), nil
	default:
		return nil, errors.New("todo")
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
		return unmarshalAsDate(data)
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

func unmarshalAsDate(data json.RawMessage) (IColumnType, error) {
	var actual DateColumnType
	var temp struct {
		Name   DataType   `json:"name"`
		Min    string     `json:"min"`
		Max    string     `json:"max"`
		Format DateFormat `json:"format"`
	}
	if err := json.Unmarshal(data, &temp); err != nil {
		return nil, err
	}

	min, err := time.Parse("2006-01-02", temp.Min)
	if err != nil {
		return nil, err
	}

	max, err := time.Parse("2006-01-02", temp.Max)
	if err != nil {
		return nil, err
	}

	actual.Name = temp.Name
	actual.Min = min
	actual.Max = max
	actual.Format = temp.Format

	fmt.Println(actual)
	return actual, nil
}
