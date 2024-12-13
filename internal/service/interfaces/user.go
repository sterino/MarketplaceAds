package interfaces

import "context"

type UserService interface {
	GetAccountTypeByID(ctx context.Context, userID string) (string, error)
}
