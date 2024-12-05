package interfaces

import "context"

type CodeRepository interface {
	SaveCode(ctx context.Context, email, code string) error
	GetCode(ctx context.Context, email string) (string, error)
	DeleteCode(ctx context.Context, email string) error
}
