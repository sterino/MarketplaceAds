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

// GetAccountTypeByID находит AccountType по ID
func (ur *UserRepository) GetAccountTypeByID(ctx context.Context, id string) (string, error) {
	var accountType string

	// Проверяем в таблице компаний
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

	// Проверяем в таблице инфлюенсеров
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

	// Если не найдено
	return "", errors.New("user not found")
}
