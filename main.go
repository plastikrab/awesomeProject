package main

import (
	"awesomeProject/handlers"
	_ "awesomeProject/handlers"
	"database/sql"
	_ "database/sql"
	"log"
	_ "log"
	"net/http"
	_ "net/http"
)

func main() {

	http.HandleFunc("/token", handlers.NewTokenHandler)
	http.HandleFunc("/startNotifications", handlers.StartNotificationHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}

	db, err := sql.Open("postgres", "user=postgres password=18082005 dbname=tokens sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
