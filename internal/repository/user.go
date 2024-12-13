package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

type UserRepository struct {
	db *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (ur *UserRepository) GetAccountTypeByID(ctx context.Context, id string) (string, error) {
	var accountType string

	queryCompany := `
		SELECT account_type
		FROM companies
		WHERE id = $1
		LIMIT 1;`
	err := ur.db.GetContext(ctx, &accountType, queryCompany, id)
	if err == nil {
		return accountType, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("error checking company: %w", err)
	}

	queryInfluencer := `
		SELECT account_type
		FROM influencer
		WHERE id = $1
		LIMIT 1;`
	err = ur.db.GetContext(ctx, &accountType, queryInfluencer, id)
	if err == nil {
		return accountType, nil
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return "", fmt.Errorf("error checking influencer: %w", err)
	}

	return "", errors.New("user not found")
}
