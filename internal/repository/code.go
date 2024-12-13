package repository

import (
	"Marketplace/internal/domain/company"
	"Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
	"log"
)

type CodeRepository struct {
	db *sqlx.DB
}

func NewCodeRepository(db *sqlx.DB) interfaces.CodeRepository {
	return &CodeRepository{
		db: db,
	}
}

func (cr *CodeRepository) SaveCode(ctx context.Context, email, code string) error {
	log.Printf(email, " ", code)
	query := `
		INSERT INTO email_verification_codes (email, code)
		VALUES ($1, $2)
		ON CONFLICT (email) DO UPDATE SET code = $2;`
	_, err := cr.db.ExecContext(ctx, query, email, code)
	return err
}

func (cr *CodeRepository) GetCode(ctx context.Context, email string) (string, error) {
	log.Printf(email)
	var code string
	query := `SELECT code FROM email_verification_codes WHERE email = $1;`
	err := cr.db.GetContext(ctx, &code, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", company.ErrorNotFound
		}
		return "", err
	}
	return code, nil
}

func (cr *CodeRepository) DeleteCode(ctx context.Context, email string) error {
	query := `DELETE FROM email_verification_codes WHERE email = $1;`
	_, err := cr.db.ExecContext(ctx, query, email)
	return err
}
