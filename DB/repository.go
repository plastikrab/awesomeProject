package DB

import (
	"awesomeProject/model/entytyes"
	_ "container/list"
	"database/sql"
	_ "github.com/claygod/coffer"
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

func (r *tokenRepository) Create(data entytyes.IncomingData) error {
	_, err := r.db.Exec("INSERT INTO Tokens (id, token) VALUES (?, ?);", data.Token, data.ID)
	return err
}

func (r *tokenRepository) GetAll() sql.Result {
	data, err := r.db.Exec("SELECT token FROM Tokens;")
	if err != nil {
		return nil
	}
	return data
}
