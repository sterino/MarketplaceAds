package interfaces

import (
	"Marketplace/internal/domain/company"
	"context"
)

type CompanyService interface {
	Register(ctx context.Context, req company.RegisterRequest) (id string, err error)
	Login(ctx context.Context, req company.LoginRequest) (string, int64, error)
	GetByEmail(ctx context.Context, email string) (company.Response, error)
	GetByEmailEntity(ctx context.Context, email string) (company.Entity, error)
	GetByID(ctx context.Context, input string) (res company.Response, err error)
	VerifyEmail(ctx context.Context, email, code string) error
}
