package interfaces

import "context"

type UserRepository interface {
	// GetAccountTypeByID возвращает тип аккаунта пользователя (company/influencer) по ID
	GetAccountTypeByID(ctx context.Context, userID string) (string, error)
}
