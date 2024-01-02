package main

import (
	"database/sql"
	"errors"
	"math/rand"
)

type TableData struct {
	Name    string    `json:"name"`
	RowData []RowData `json:"rows"`
}

func GenerateTableData(desc TableDesc, dataTables []TableData, r *rand.Rand, db *sql.DB) (TableData, error) {
	t := TableData{
		Name:    desc.Name,
		RowData: []RowData{},
	}

	// We track the number of iterations that have occurred -- the number of attempts to add a row to the table. If the
	// number of iterations exceeds 3x the number of rows (only 1/3 iterations inserted a row) then we assume that it
	// is either not possible to fulfill the request or not worthwile with the current approach.
	//
	// Developing a method that produces unique random values for unique/primary key columns would be a useful
	// improvement
	it := uint(0)
	row := uint(0)
	for row < desc.NumRows {
		if it > 3*desc.NumRows {
			return t, errors.New("exceeded maximum attempts")
		}

		data := RowData{
			Entries: make(map[string]RowEntry),
		}

		for _, columnDesc := range desc.Columns {
			if entry, err := columnDesc.Type.GenerateEntry(desc, row, r, db); err != nil {
				return t, err
			} else {
				data.Entries[columnDesc.Name] = entry
			}
		}

		for _, referenceDesc := range desc.References {
			if entry, err := referenceDesc.GenerateEntry(dataTables, t, r); err != nil {
				return t, err
			} else {
				data.Entries[referenceDesc.Name] = entry
			}
		}

		if !data.HasOverlapConflict(desc, t.RowData) {
			t.RowData = append(t.RowData, data)
			row += 1
		}

		it += 1
	}

	return t, nil
}

func GenerateTables(descs []TableDesc, r *rand.Rand, db *sql.DB) ([]TableData, error) {
	var dataTables []TableData

	for _, table := range descs {
		t, err := GenerateTableData(table, dataTables, r, db)
		if err != nil {
			return dataTables, err
		}

		dataTables = append(dataTables, t)
	}

	return dataTables, nil
}
