package main

import (
	"awesomeProject/handlers"
	"database/sql"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/token", handlers.NewTokenHandler)
	http.ListenAndServe(":8080", nil)

	db, err := sql.Open("postgres", "user=postgres password=18082005 dbname=tokens sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
