package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("chnge")
	_, err := fmt.Fprint(w, "Hello, World!")
	if err != nil {
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
