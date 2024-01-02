package main

import "encoding/json"

type RowEntry interface{}

type RowData struct {
	Entries map[string]RowEntry
}

func (rd RowData) MarshalJSON() ([]byte, error) {
	return json.Marshal(rd.Entries)
}

func (rd RowData) HasOverlapConflict(desc TableDesc, rows []RowData) bool {
	var hasPrimaryKey bool
	if desc.numPrimaryKeys() > 0 {
		hasPrimaryKey = true
	} else {
		hasPrimaryKey = false
	}

	for _, row := range rows {
		anyKeyDistinct := false
		for _, col := range desc.Columns {
			if row.Entries[col.Name] != rd.Entries[col.Name] {
				if col.PrimaryKey {
					anyKeyDistinct = true
				}
			} else if col.Unique {
				return true
			}
		}

		for _, ref := range desc.References {
			if row.Entries[ref.Name] != rd.Entries[ref.Name] {
				if ref.PrimaryKey {
					anyKeyDistinct = true
				}
			} else if ref.Unique {
				return true
			}
		}

		if hasPrimaryKey && !anyKeyDistinct {
			return true
		}
	}

	return false
}
