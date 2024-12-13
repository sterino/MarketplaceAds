package interfaces

import (
	"Marketplace/internal/domain/application"
	"context"
)

type ApplicationService interface {
	Create(ctx context.Context, req application.CreateRequest) (application.Response, error)
	GetByID(ctx context.Context, id string) (application.Response, error)
	GetByAdID(ctx context.Context, adID string) ([]application.Response, error)
	GetByInfluencerID(ctx context.Context, influencerID string) ([]application.Response, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	Delete(ctx context.Context, id string) error
}
