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
		FROM api_token 
		WHERE api_token = ?`, result.Token).
		Scan(
			&result.ID,
			&result.AccountID,
			&result.CreatedAt,
			&result.DeletedAt,
		)

	if err == gocql.ErrNotFound {
		err = api.ErrNotFound
	}

	return &result, err
}

func (t *TokenRepo) CreateToken(accountID string) (*model.Token, error) {
	token := &model.Token{
		ID:        GenerateUUID(),
		AccountID: accountID,
		Token:     GenerateRandomString(32),
	}

	return token, t.db.Query(`
		INSERT INTO api_token(
			id,
			account_id,
			api_token,
			created_at,
			deleted_at
		) VALUES(?,?,?,DATEOF(NOW()),'')`,
		token.ID,
		token.AccountID,
		token.Token).Exec()
}

func (t *TokenRepo) DeleteToken(tokenID string, accountID string) error {
	return t.db.Query(`
		UPDATE api_token
		SET 
			deleted_at = DATEOF(NOW())
			api_token = ?, 
		WHERE
			id = ?
		AND account_id = ?`, GenerateZeros(32), tokenID, accountID).Exec()
}
