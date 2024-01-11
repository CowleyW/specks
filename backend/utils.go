package main

import (
	"archive/zip"
	"bytes"
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

func zipFiles(files []OutputFile) ([]byte, error) {
	buf := new(bytes.Buffer)
	zipWriter := zip.NewWriter(buf)

	for _, file := range files {
		zipFile, err := zipWriter.Create(file.Name)
		if err != nil {
			return nil, err
		}
		_, err = zipFile.Write([]byte(file.Data))
		if err != nil {
			return nil, err
		}
	}

	err := zipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
