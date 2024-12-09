package handlers

import (
	"awesomeProject/DB"
	"awesomeProject/apis"
	"awesomeProject/model/entytyes"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
)

func StartNotificationHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	var token entytyes.TokenData
	err := json.NewDecoder(r.Body).Decode(&token)

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

	tokensList, err := repo.GetAll()
	if err != nil {
		log.Println(err)
	}
	if len(tokensList) == 0 || tokensList == nil {
		log.Println("Пустой список токенов")
		return
	}
	registrationTokens := []string{}
	for _, t := range tokensList {
		registrationTokens = append(registrationTokens, t.Token)
	}
	err = apis.NotifyFirebase(registrationTokens, "Notificatin", "Description")
	if err != nil {
		return
	}
}
