package repository

import (
	"Marketplace/internal/domain/influencer"
	interfaces "Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
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
	influencerId := uuid.New().String()
	query := `
		INSERT INTO influencers (id, name, email, password, phone_number, platforms, followers_count, category, bio, address)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id;`
	args := []any{
		influencerId,
		data.Name,
		data.Email,
		data.Password,
		data.PhoneNumber,
		data.Platforms,
		data.FollowersCount,
		data.Category,
		data.Bio,
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
		SELECT id, name, email, password, phone_number, platforms, followers_count, category, bio, address, created_at, updated_at
		FROM influencers
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
		SELECT id, name, email, password, phone_number, platforms, followers_count, category, bio, address, created_at, updated_at
		FROM influencers
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

func (ir *InfluencerRepository) UpdateEmailVerification(ctx context.Context, id string) error {
	query := `UPDATE ads SET email_verified = $1, updated_at = NOW() WHERE id = $2;`

	_, err := ir.db.ExecContext(ctx, query, true, id)
	return err
}
