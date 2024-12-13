package repository

import (
	"Marketplace/internal/domain/company"
	interfaces "Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type CompanyRepository struct {
	db *sqlx.DB
}

func NewCompanyRepository(db *sqlx.DB) interfaces.CompanyRepository {
	return &CompanyRepository{
		db: db,
	}
}

func (cr *CompanyRepository) Create(ctx context.Context, data company.RegisterRequest) (id string, err error) {
	query := `
		INSERT INTO companies (name, email, password, phone_number, address)
		VALUES ($1, $2, $3, $4, $5) RETURNING id;`
	args := []any{
		data.Name,
		data.Email,
		data.Password,
		data.PhoneNumber,
		data.Address,
	}
	if err = cr.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = company.ErrorNotFound
		}
	}
	return
}

func (cr *CompanyRepository) GetByEmail(ctx context.Context, email string) (dest company.Entity, err error) {
	query := `
		SELECT * FROM companies WHERE email = $1;`
	err = cr.db.GetContext(ctx, &dest, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = company.ErrorNotFound
		}
		return
	}
	return
}

func (cr *CompanyRepository) GetByID(ctx context.Context, id string) (dest company.Entity, err error) {
	query := `
		SELECT * FROM companies WHERE id = $1;`
	err = cr.db.GetContext(ctx, &dest, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = company.ErrorNotFound
		}
		return
	}
	return
}

func (cr *CompanyRepository) UpdateEmailVerification(ctx context.Context, id string) error {
	query := `UPDATE ads SET email_verified = $1, updated_at = NOW() WHERE id = $2;`

	_, err := cr.db.ExecContext(ctx, query, true, id)
	return err
}
