package interfaces

import (
	"Marketplace/internal/domain/order"
	"context"
)

type OrderRepository interface {
	Create(ctx context.Context, entity order.Entity) (string, error)
	GetByID(ctx context.Context, id string) (order.Entity, error)
	UpdateStatus(ctx context.Context, id, status string) error
	GetByCompanyID(ctx context.Context, companyID string) ([]order.Entity, error)
	GetByInfluencerID(ctx context.Context, influencerID string) ([]order.Entity, error)
	Delete(ctx context.Context, id string) error
}
