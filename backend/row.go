package main

import "encoding/json"

type RowEntry interface{}

type RowData struct {
	Entries []RowEntry
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
		for i, col := range desc.Columns {
			if row.Entries[i] != rd.Entries[i] {
				if col.PrimaryKey {
					anyKeyDistinct = true
				}
			} else if col.Unique {
				return true
			}
		}

		for i, ref := range desc.References {
			if row.Entries[i] != rd.Entries[i] {
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
