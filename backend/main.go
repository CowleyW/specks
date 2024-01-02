package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
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
	db, err := openDB("backend:password@tcp(db-dev:3306)/specks_db")
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

	http.HandleFunc("/", app.handler)
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
