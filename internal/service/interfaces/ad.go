package interfaces

import (
	"Marketplace/internal/domain/ad"
	"context"
)

type AdService interface {
	Create(ctx context.Context, input ad.CreateRequest) (string, error)

	GetByID(ctx context.Context, id string) (ad.Response, error)

	GetAll(ctx context.Context) ([]ad.Response, error)

	UpdateStatus(ctx context.Context, id string, status string) error

	GetByCompanyID(ctx context.Context, companyID string) ([]ad.Response, error)

	Delete(ctx context.Context, id string) error
}
