package interfaces

import (
	"Marketplace/internal/domain/order"
	"context"
)

type OrderService interface {
	Create(ctx context.Context, req order.CreateRequest) (order.Response, error)
	GetByID(ctx context.Context, id string) (order.Response, error)
	UpdateStatus(ctx context.Context, id string, status string) error
	GetByCompanyID(ctx context.Context, companyID string) ([]order.Response, error)
	GetByInfluencerID(ctx context.Context, influencerID string) ([]order.Response, error)
	Delete(ctx context.Context, id string) error
}
