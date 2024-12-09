package DB

import (
	"awesomeProject/model/entytyes"
	_ "container/list"
	"database/sql"
	_ "github.com/claygod/coffer"
	"log"
)

type TokensRepository interface {
	Create(token string)
	GetAll() []string
}

type tokenRepository struct {
	db *sql.DB
}

func NewTokenRepository(db *sql.DB) *tokenRepository {
	return &tokenRepository{db: db}
}

func (r *tokenRepository) Create(data entytyes.TokenData) error {
	println("Start database query")
	_, err := r.db.Exec("INSERT INTO savedTokens (tokens) VALUES ($1);", data.Token)
	return err
}

func (r *tokenRepository) GetAll() ([]entytyes.TokenData, error) {
	log.Println("Start database query")
	var tokensList []entytyes.TokenData
	rows, err := r.db.Query("SELECT * FROM savedTokens")
	if err != nil {
		return nil, err // Return the error instead of panicking
	}
	defer rows.Close()

	for rows.Next() {
		var token entytyes.TokenData
		err := rows.Scan(&token.Token)
		if err != nil {
			return nil, err
		}
		tokensList = append(tokensList, token)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tokensList, nil
}
