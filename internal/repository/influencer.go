package repository

import (
	"Marketplace/internal/domain/influencer"
	interfaces "Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type InfluencerRepository struct {
	db *sqlx.DB
}

func NewInfluencerRepository(db *sqlx.DB) interfaces.InfluencerRepository {
	return &InfluencerRepository{
		db: db,
	}
}

func (ir *InfluencerRepository) Create(ctx context.Context, data influencer.RegisterRequest) (id string, err error) {
	query := `
		INSERT INTO influencer (name, email, password, phone_number, platforms, followers_count, category, bio, price_per_post, address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;`
	args := []any{
		data.Name,
		data.Email,
		data.Password,
		data.PhoneNumber,
		data.Platforms, // массив ссылок
		data.FollowersCount,
		data.Category,
		data.Bio,
		data.PricePerPost,
		data.Address,
	}
	if err = ir.db.QueryRowContext(ctx, query, args...).Scan(&id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = influencer.ErrorNotFound
		}
	}
	return
}

func (ir *InfluencerRepository) GetByEmail(ctx context.Context, email string) (dest influencer.Entity, err error) {
	query := `
		SELECT id, name, email, password, phone_number, platforms, followers_count, category, bio, price_per_post, address, created_at, updated_at
		FROM influencer
		WHERE email = $1;`
	err = ir.db.GetContext(ctx, &dest, query, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = influencer.ErrorNotFound
		}
		return
	}
	return
}

func (ir *InfluencerRepository) GetByID(ctx context.Context, id string) (dest influencer.Entity, err error) {
	query := `
		SELECT id, name, email, password, phone_number, platforms, followers_count, category, bio, price_per_post, address, created_at, updated_at
		FROM influencer
		WHERE id = $1;`
	err = ir.db.GetContext(ctx, &dest, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err = influencer.ErrorNotFound
		}
		return
	}
	return
}
