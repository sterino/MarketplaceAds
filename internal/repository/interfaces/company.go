package interfaces

import (
	"Marketplace/internal/domain/company"
	"context"
)

type CompanyRepository interface {
	Create(ctx context.Context, data company.RegisterRequest) (id string, err error)
	GetByEmail(ctx context.Context, email string) (data company.Entity, err error)
	GetByID(ctx context.Context, id string) (data company.Entity, err error)
}
