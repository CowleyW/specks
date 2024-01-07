package main

import (
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

	var descs []TableDesc
	if err := json.NewDecoder(r.Body).Decode(&descs); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	fmt.Printf("Received: %+v\n", descs)

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataTables, err := GenerateTables(descs, random, app.db)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to generate data", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(dataTables)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error constructing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "error building the response", http.StatusInternalServerError)
		return
	}
}

func (app application) previewHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var descs []TableDesc
	if err := json.NewDecoder(r.Body).Decode(&descs); err != nil {
		fmt.Println(err)
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	for i := 0; i < len(descs); i += 1 {
		descs[i].NumRows = 10
	}

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	dataTables, err := GenerateTables(descs, random, app.db)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "failed to generate data", http.StatusBadRequest)
		return
	}

	jsonData, err := json.Marshal(dataTables)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Error constructing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(jsonData)
	if err != nil {
		http.Error(w, "error building the response", http.StatusInternalServerError)
		return
	}
}
