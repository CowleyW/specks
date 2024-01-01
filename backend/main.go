package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func handler(w http.ResponseWriter, r *http.Request) {
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
	dataTables, err := GenerateTables(descs, random)
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

func addCORSHeaders(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from all origins
		w.Header().Set("Access-Control-Allow-Origin", "*")

		// Allow certain headers in the request
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Allow specific methods (POST in this case)
		w.Header().Set("Access-Control-Allow-Methods", "POST")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		handler.ServeHTTP(w, r)
	})
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", addCORSHeaders(http.DefaultServeMux)))
}
