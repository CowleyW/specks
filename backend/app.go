package main

import (
	"archive/zip"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"
)

type application struct {
	db *sql.DB
}

func (app application) generateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	app.generateDataWithLimit(w, r, 1000, 20)
}

func (app application) previewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	app.generateDataWithLimit(w, r, 10, 1)
}

func (app application) generateDataWithLimit(w http.ResponseWriter, r *http.Request, rowLimit, tableLimit uint) {
	var spec DescriptionSpec
	if err := json.NewDecoder(r.Body).Decode(&spec); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(spec.TableDescs); i += 1 {
		if spec.TableDescs[i].NumRows > rowLimit {
			spec.TableDescs[i].NumRows = rowLimit
		}
	}

	if len(spec.TableDescs) > int(tableLimit) {
		spec.TableDescs = spec.TableDescs[0:tableLimit]
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataTables, err := GenerateTables(spec.TableDescs, random, app.db)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to generate data", http.StatusBadRequest)
		return
	}

	var outputData []byte
	switch spec.OutputFormat {
	case CSV:
		var csvFiles []OutputFile
		for _, table := range dataTables {
			csvFiles = append(csvFiles, FormatAsCSV(table))
		}

		if spec.ForPreview {
			outputData = []byte(csvFiles[0].Data)
			w.Header().Set("Content-Type", "text/csv")
		} else {
			buf := new(bytes.Buffer)
			zipWriter := zip.NewWriter(buf)

			for _, file := range csvFiles {
				zipFile, err := zipWriter.Create(file.Name)
				if err != nil {
					fmt.Println(err)
					http.Error(w, "Failed to generate data", http.StatusInternalServerError)
					return
				}
				_, err = zipFile.Write([]byte(file.Data))
				if err != nil {
					fmt.Println(err)
					http.Error(w, "Failed to generate data", http.StatusInternalServerError)
					return
				}
			}

			err := zipWriter.Close()
			if err != nil {
				fmt.Println(err)
				http.Error(w, "Failed to generate data", http.StatusInternalServerError)
				return
			}

			outputData = buf.Bytes()
			w.Header().Set("Content-Type", "application/zip")
		}
	case JSON:
		outputData, err = json.Marshal(dataTables)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "Error constructing data", http.StatusInternalServerError)
			return
		}

		if spec.ForPreview {
			w.Header().Set("Content-Type", "application/json")
		} else {
			http.Error(w, "Not yet implemented", http.StatusInternalServerError)
		}
	case SQL:
	}

	w.WriteHeader(http.StatusOK)
	_, err = w.Write(outputData)
	if err != nil {
		http.Error(w, "error building the response", http.StatusInternalServerError)
		return
	}
}
