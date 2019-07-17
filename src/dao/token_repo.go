package dao

import (
	"github.com/gocql/gocql"

	"github.com/kachan1208/auth/src/api"
	"github.com/kachan1208/auth/src/model"
)

type TokenRepo struct {
	db *gocql.Session
}

func NewTokenRepo(db *gocql.Session) *TokenRepo {
	return &TokenRepo{
		db: db,
	}
}

func (t *TokenRepo) GetToken(token string) (*model.Token, error) {
	result := model.Token{
		Token: token,
	}

	err := t.db.Query(`
		SELECT 
			id, 
			account_id, 
			created_at, 
			deleted_at 
		FROM token 
		WHERE token = ?`, result.Token).
		Scan(
			&result.ID,
			&result.AccountID,
			&result.CreatedAt,
			&result.DeletedAt,
		)

	if err == gocql.ErrNotFound {
		err = api.ErrNotFound
	}
	return result, err
}

func (t *TokenRepo) CreateToken(token *model.Token) error {
	token.ID = GenerateUUID()
	token.Token = GenerateRandomString(64)

	err := t.db.Query(`
		INSERT INTO token(
			id,
			account_id,
			created_at,
			deleted_at,
		) VALUES(?,?,DATEOF(NOW()),'') 
		`, token.Token).
		Scan(
			&token.ID,
			&token.AccountID,
			&token.CreatedAt,
			&token.DeletedAt,
		)

	return err
}
