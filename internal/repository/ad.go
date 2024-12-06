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

// Create добавляет новое объявление в базу данных
func (ar *AdRepository) Create(ctx context.Context, data ad.CreateRequest) (string, error) {
	// Валидация данных перед вставкой
	if err := data.Validate(); err != nil {
		return "", err
	}

	// Запрос на создание объявления
	query := `
		INSERT INTO ads (title, description, price, status)
		VALUES ($1, $2, $3, $4)
		RETURNING id;`
	args := []interface{}{data.Title, data.Description, data.Price, "open"}

	var id string
	err := ar.db.QueryRowContext(ctx, query, args...).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

// GetByID возвращает объявление по его ID
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

// GetAll возвращает все объявления
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

// UpdateStatus обновляет статус объявления
func (ar *AdRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE ads SET status = $1, updated_at = NOW() WHERE id = $2;`

	_, err := ar.db.ExecContext(ctx, query, status, id)
	return err
}
