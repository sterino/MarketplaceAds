package interfaces

import (
	"Marketplace/internal/domain/influencer"
	"context"
)

type InfluencerRepository interface {
	Create(ctx context.Context, data influencer.RegisterRequest) (id string, err error)
	GetByEmail(ctx context.Context, email string) (data influencer.Entity, err error)
	GetByID(ctx context.Context, id string) (data influencer.Entity, err error)
}
