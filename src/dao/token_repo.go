package dao

import (
	"time"

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

func (t *TokenRepo) GetTokenByToken(token string) (*model.Token, error) {
	result := model.Token{
		Token: token,
	}

	err := t.db.Query(`
		SELECT 
			id,
			account_id,
			is_enabled,
			created_at, 
			deleted_at 
		FROM api_token 
		WHERE api_token = ?`, result.Token).
		Scan(
			&result.ID,
			&result.AccountID,
			&result.IsEnabled,
			&result.CreatedAt,
			&result.DeletedAt,
		)

	if err == gocql.ErrNotFound || result.IsDeleted() {
		err = api.ErrNotFound
	}

	return &result, err
}

func (t *TokenRepo) GetTokenByID(id string) (*model.Token, error) {
	result := model.Token{}

	err := t.db.Query(`
		SELECT 
			id,
			account_id,
			is_enabled,
			created_at, 
			deleted_at 
		FROM api_token 
		WHERE id = ?`, id).
		Scan(
			&result.ID,
			&result.AccountID,
			&result.IsEnabled,
			&result.CreatedAt,
			&result.DeletedAt,
		)

	if err == gocql.ErrNotFound || result.IsDeleted() {
		err = api.ErrNotFound
	}

	return &result, err
}

func (t *TokenRepo) CreateToken(accountID string) (*model.Token, error) {
	token := &model.Token{
		ID:        GenerateUUID(),
		AccountID: accountID,
		IsEnabled: true,
		Token:     GenerateRandomString(32),
	}

	return token, t.db.Query(`
		INSERT INTO api_token(
			id,
			account_id,
			api_token,
			is_enabled,
			created_at,
			deleted_at
		) VALUES(?,?,?,DATEOF(NOW()),'')`,
		token.ID,
		token.AccountID,
		token.IsEnabled,
		token.Token).Exec()
}

func (t *TokenRepo) DeleteToken(id string, accountID string) error {
	return t.db.Query(`
		UPDATE api_token
		SET 
			deleted_at = DATEOF(NOW())
			api_token = ?, 
		WHERE
			id = ?
		AND account_id = ?`, GenerateZeros(32), id, accountID).Exec()
}

func (t *TokenRepo) UpdateToken(id string, accountID string, isEnabled bool) error {
	return t.db.Query(`
		UPDATE api_token
		SET 
			is_enabled = ? 
		WHERE
			id = ?
		AND account_id = ?`, isEnabled, id, accountID).Exec()
}

func (t *TokenRepo) TokenList(accountID string) ([]model.Token, error) {
	iter := t.db.Query(`
		SELECT 
			id,
			is_enabled,
			created_at, 
			deleted_at 
		FROM api_token 
		WHERE account_id = ?`, accountID).
		Iter()

	var (
		id        string
		isEnabled bool
		createdAt time.Time
		deletedAt time.Time
	)
	result := make([]model.Token, 0, 0)
	for iter.Scan(&id, &isEnabled, &createdAt, &deletedAt) {
		result = append(result, model.Token{
			ID:        id,
			IsEnabled: isEnabled,
			CreatedAt: createdAt,
			DeletedAt: deletedAt,
		})
	}

	err := iter.Close()
	if err == gocql.ErrNotFound {
		err = api.ErrNotFound
	}

	return result, err
}
