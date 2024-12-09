package handlers

import (
	"awesomeProject/DB"
	"awesomeProject/model/entytyes"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func NewTokenHandler(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("postgres", "user=postgres password=18082005 dbname=tokens sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	} else {
		log.Println("Successfully Connected")
	}
	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(db)
	repo := DB.NewTokenRepository(db)
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var token entytyes.TokenData
	err = json.NewDecoder(r.Body).Decode(&token)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	err = repo.Create(token)
	if err != nil {
		println("Error: " + err.Error())
		return
	}
	fmt.Println(token.Token)
	fmt.Fprintf(w, "Goyda")
}
