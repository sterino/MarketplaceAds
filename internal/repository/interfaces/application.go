package interfaces

import (
	"Marketplace/internal/domain/application"
	"context"
)

type ApplicationRepository interface {
	Create(ctx context.Context, entity application.Entity) (string, error)
	GetByID(ctx context.Context, id string) (application.Entity, error)
	GetByAdID(ctx context.Context, adID string) ([]application.Entity, error)
	GetByInfluencerID(ctx context.Context, influencerID string) ([]application.Entity, error)
	UpdateStatus(ctx context.Context, id, status string) error
	Delete(ctx context.Context, id string) error
}
