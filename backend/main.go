package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
	"os"
)

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
	pw, err := readSecret("db_pw")
	if err != nil {
		log.Fatalln(err)
	}

	dbAddr := os.Getenv("DATABASE_ADDRESS")

	db, err := openDB(fmt.Sprintf("backend:%s@tcp(%s)/specks_db", pw, dbAddr))
	if err != nil {
		log.Fatalln(err)
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalln(err)
		}
	}(db)

	app := application{db: db}

	http.HandleFunc("/", app.generateHandler)
	http.HandleFunc("/preview/", app.previewHandler)
	log.Fatal(http.ListenAndServe(":4000", addCORSHeaders(http.DefaultServeMux)))
}

func openDB(dns string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dns)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func readSecret(filename string) (string, error) {
	content, err := ioutil.ReadFile(fmt.Sprintf("/run/secrets/%s", filename))
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(content)), nil
}
