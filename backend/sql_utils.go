package main

import (
	"database/sql"
	"errors"
	"fmt"
	"math/rand"
)

func queryRandomColumn(db *sql.DB, r *rand.Rand, tableName, columnName string) (string, error) {
	var count int
	err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", tableName)).Scan(&count)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("failed to scan query result for count")
	}

	random := r.Intn(count) + 1
	var name string
	err = db.QueryRow(fmt.Sprintf("SELECT %s FROM %s WHERE id = ?", columnName, tableName), random).Scan(&name)
	if err != nil {
		fmt.Println("Random number: ", random)
		return "", errors.New("failed to scan query result for name")
	}

	return name, nil
}
