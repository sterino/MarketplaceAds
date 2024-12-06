package interfaces

import "context"

type UserService interface {
	// GetAccountTypeByID получает тип аккаунта пользователя (company/influencer) по ID
	GetAccountTypeByID(ctx context.Context, userID string) (string, error)
}
