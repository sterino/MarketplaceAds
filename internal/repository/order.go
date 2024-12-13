package repository

import (
	"Marketplace/internal/domain/order"
	"context"
	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	db *sqlx.DB
}

func NewOrderRepository(db *sqlx.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (r *OrderRepository) Create(ctx context.Context, entity order.Entity) (string, error) {
	query := `INSERT INTO orders (ad_id, company_id, influencer_id, status, price, description, created_at, updated_at)
              VALUES ($1, $2, $3, $4, $5, $6, NOW(), NOW()) RETURNING id`
	var id string
	err := r.db.QueryRowContext(ctx, query, entity.AdID, entity.CompanyID, entity.InfluencerID, entity.Status, entity.Price, entity.Description).Scan(&id)
	return id, err
}

func (r *OrderRepository) GetByID(ctx context.Context, id string) (order.Entity, error) {
	query := `SELECT * FROM orders WHERE id = $1`
	var entity order.Entity
	err := r.db.GetContext(ctx, &entity, query, id)
	return entity, err
}

func (r *OrderRepository) UpdateStatus(ctx context.Context, id, status string) error {
	query := `UPDATE orders SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *OrderRepository) GetByCompanyID(ctx context.Context, companyID string) ([]order.Entity, error) {
	query := `SELECT * FROM orders WHERE company_id = $1`
	var entities []order.Entity
	err := r.db.SelectContext(ctx, &entities, query, companyID)
	return entities, err
}

func (r *OrderRepository) GetByInfluencerID(ctx context.Context, influencerID string) ([]order.Entity, error) {
	query := `SELECT * FROM orders WHERE influencer_id = $1`
	var entities []order.Entity
	err := r.db.SelectContext(ctx, &entities, query, influencerID)
	return entities, err
}

func (r *OrderRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM orders WHERE id = $1`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
