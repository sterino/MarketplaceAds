package repository

import (
	"Marketplace/internal/domain/ad"
	"Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type AdRepository struct {
	db *sqlx.DB
}

func NewAdRepository(db *sqlx.DB) interfaces.AdRepository {
	return &AdRepository{db: db}
}

func (ar *AdRepository) Create(ctx context.Context, data ad.Entity) (string, error) {

	if err := data.Validate(); err != nil {
		return "", err
	}

	query := `
		INSERT INTO ads (id, title, description, price, status, company_id)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id;`
	args := []interface{}{data.Title, data.Description, data.Price, "open", data.CompanyID}

	var id string
	err := ar.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ar *AdRepository) GetByID(ctx context.Context, id string) (ad.Entity, error) {
	query := `SELECT id, title, description, price, status, orders, created_at, updated_at 
		FROM ads WHERE id = $1;`
	var entity ad.Entity

	err := ar.db.GetContext(ctx, &entity, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return ad.Entity{}, ad.ErrorNotFound
	}
	return entity, err
}

func (ar *AdRepository) GetAll(ctx context.Context) ([]ad.Entity, error) {
	query := `SELECT id, title, description, price, status, orders, created_at, updated_at 
		FROM ads;`
	var ads []ad.Entity

	err := ar.db.SelectContext(ctx, &ads, query)
	if err != nil {
		return nil, err
	}
	return ads, nil
}

func (ar *AdRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE ads SET status = $1, updated_at = NOW() WHERE id = $2;`

	_, err := ar.db.ExecContext(ctx, query, status, id)
	return err
}

func (ar *AdRepository) GetByCompanyID(ctx context.Context, companyID string) ([]ad.Entity, error) {
	query := `SELECT id, title, description, price, status, orders_id, created_at, updated_at 
		FROM ads WHERE company_id = $1;`
	var ads []ad.Entity

	err := ar.db.SelectContext(ctx, &ads, query, companyID)
	return ads, err
}

func (ar *AdRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM ads WHERE id = $1;`
	_, err := ar.db.ExecContext(ctx, query, id)
	return err
}
