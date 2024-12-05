package service

import (
	"Marketplace/internal/domain/company"
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/jwt"
	"Marketplace/internal/utils/password"
	"context"
	"errors"
	"time"
)

type CompanyService struct {
	companyRepository interfaces.CompanyRepository
	secretKey         []byte
}

func NewCompanyService(companyRepository interfaces.CompanyRepository, secretKey []byte) services.CompanyService {
	return &CompanyService{
		companyRepository: companyRepository,
		secretKey:         secretKey,
	}
}

func (s *CompanyService) Login(ctx context.Context, input company.LoginRequest) (string, int64, error) {
	user, err := s.GetByEmailEntity(ctx, input.Email)
	if err != nil {
		return "", 0, err
	}

	ok := password.Compare(input.Password, user.Password)
	if !ok {
		return "", 0, errors.New("invalid credentials")
	}

	token, expiresAt, err := jwt.Encode(jwt.JWT{
		UUID:  user.ID,
		Email: user.Email,
	}, s.secretKey)

	if err != nil {
		return "", 0, err
	}

	return *token, *expiresAt, nil

}

func (s *CompanyService) Register(ctx context.Context, input company.RegisterRequest) (string, error) {
	_, err := s.companyRepository.GetByEmail(ctx, input.Email)
	if err == nil {
		return "", company.ErrorEmailConflict
	} else if !errors.Is(err, company.ErrorNotFound) {
		return "", err
	}

	hashedPassword, err := password.Generate(input.Password)
	if err != nil {
		return "", err
	}

	newCompany := company.RegisterRequest{
		Name:        input.Name,
		Email:       input.Email,
		Password:    hashedPassword,
		PhoneNumber: input.PhoneNumber,
		Address:     input.Address,
	}
	return s.Create(ctx, newCompany)
}

func (s *CompanyService) GetByEmail(ctx context.Context, input string) (res company.Response, err error) {
	user, err := s.companyRepository.GetByEmail(ctx, input)

	if err != nil {
		return
	}
	res = company.ParseFromEntity(user)
	return
}

func (s *CompanyService) GetByEmailEntity(ctx context.Context, input string) (user company.Entity, err error) {
	user, err = s.companyRepository.GetByEmail(ctx, input)

	if err != nil {
		return
	}
	return
}

func (s *CompanyService) GetByID(ctx context.Context, input string) (res company.Response, err error) {
	user, err := s.companyRepository.GetByID(ctx, input)

	if err != nil {
		return
	}
	res = company.ParseFromEntity(user)
	return
}

func (s *CompanyService) Create(ctx context.Context, data company.RegisterRequest) (id string, err error) {
	id, err = s.companyRepository.Create(ctx, company.RegisterRequest{
		Name:        data.Name,
		Email:       data.Email,
		Password:    data.Password,
		PhoneNumber: data.PhoneNumber,
		Address:     data.Address,
	})
	if err != nil {
		return
	}

	return
}

func (s *CompanyService) VerifyCode(ctx context.Context, email, code string) error {
	// Получаем сохраненный код и время истечения
	storedCode, expiresAt, err := s.companyRepository.GetVerificationCode(ctx, email)
	if err != nil {
		return err
	}

	// Проверка срока действия кода
	if time.Now().After(expiresAt) {
		return errors.New("verification code has expired")
	}

	// Проверка кода
	if storedCode != code {
		return errors.New("invalid verification code")
	}

	// Успешная проверка — обновляем статус пользователя
	err = s.companyRepository.MarkEmailVerified(ctx, email)
	if err != nil {
		return err
	}

	return nil
}
