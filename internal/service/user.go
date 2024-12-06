package service

import (
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"context"
	"database/sql"
	"errors"
)

type UserService struct {
	userRepository interfaces.UserRepository
}

func NewUserService(userRepository interfaces.UserRepository) services.UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// GetAccountTypeByID возвращает тип аккаунта пользователя (company/influencer) по его ID
func (s *UserService) GetAccountTypeByID(ctx context.Context, userID string) (string, error) {
	accountType, err := s.userRepository.GetAccountTypeByID(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", errors.New("user not found")
		}
		return "", err
	}
	return accountType, nil
}
