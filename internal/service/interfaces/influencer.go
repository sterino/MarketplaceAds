package interfaces

import (
	"Marketplace/internal/domain/influencer"
	"context"
)

type InfluencerService interface {
	Register(ctx context.Context, req influencer.RegisterRequest) (id string, err error)
	Login(ctx context.Context, req influencer.LoginRequest) (string, int64, error)
	GetByEmail(ctx context.Context, email string) (influencer.Response, error)
	GetByEmailEntity(ctx context.Context, email string) (influencer.Entity, error)
	GetByID(ctx context.Context, input string) (res influencer.Response, err error)
}
