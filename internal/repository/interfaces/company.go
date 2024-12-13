package interfaces

import (
	"Marketplace/internal/domain/company"
	"context"
)

type CompanyRepository interface {
	Create(ctx context.Context, data company.Entity) (id string, err error)
	GetByEmail(ctx context.Context, email string) (data company.Entity, err error)
	GetByID(ctx context.Context, id string) (data company.Entity, err error)
	UpdateEmailVerification(ctx context.Context, id string) error
}
