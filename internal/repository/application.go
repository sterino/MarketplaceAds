package repository

import (
	"Marketplace/internal/domain/application"
	"Marketplace/internal/repository/interfaces"
	"context"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type ApplicationRepository struct {
	db *sqlx.DB
}

func NewApplicationRepository(db *sqlx.DB) interfaces.ApplicationRepository {
	return &ApplicationRepository{db: db}
}

func (r *ApplicationRepository) Create(ctx context.Context, data application.Entity) (string, error) {
	query := `
		INSERT INTO applications (ad_id, company_id, influencer_id)
		VALUES ($1, $2, $3)
		RETURNING id;
	`
	var id string
	err := r.db.QueryRowContext(ctx, query, data.AdID, data.CompanyID, data.InfluencerID).Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r *ApplicationRepository) GetByID(ctx context.Context, id string) (application.Entity, error) {
	query := `SELECT id, ad_id, company_id, influencer_id, status FROM applications WHERE id = $1;`
	var app application.Entity
	err := r.db.GetContext(ctx, &app, query, id)
	if errors.Is(err, sql.ErrNoRows) {
		return app, application.ErrorNotFound
	}
	return app, err
}

func (r *ApplicationRepository) GetByAdID(ctx context.Context, adID string) ([]application.Entity, error) {
	query := `SELECT id, ad_id, company_id, influencer_id, status FROM applications WHERE ad_id = $1;`
	var apps []application.Entity
	err := r.db.SelectContext(ctx, &apps, query, adID)
	return apps, err
}

func (r *ApplicationRepository) GetByInfluencerID(ctx context.Context, influencerID string) ([]application.Entity, error) {
	query := `SELECT id, ad_id, company_id, influencer_id, status FROM applications WHERE influencer_id = $1;`
	var apps []application.Entity
	err := r.db.SelectContext(ctx, &apps, query, influencerID)
	return apps, err
}

func (r *ApplicationRepository) UpdateStatus(ctx context.Context, id string, status string) error {
	query := `UPDATE applications SET status = $1 WHERE id = $2;`
	_, err := r.db.ExecContext(ctx, query, status, id)
	return err
}

func (r *ApplicationRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM applications WHERE id = $1;`
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
