package service

import (
	"Marketplace/internal/domain/influencer"
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"Marketplace/internal/utils/jwt"
	"Marketplace/internal/utils/password"
	"context"
	"errors"
)

type InfluencerService struct {
	influencerRepository interfaces.InfluencerRepository
	secretKey            []byte
}

func NewInfluencerService(influencerRepository interfaces.InfluencerRepository, secretKey []byte) services.InfluencerService {
	return &InfluencerService{
		influencerRepository: influencerRepository,
		secretKey:            secretKey,
	}
}

func (s *InfluencerService) Login(ctx context.Context, input influencer.LoginRequest) (string, int64, error) {
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

func (s *InfluencerService) Register(ctx context.Context, input influencer.RegisterRequest) (string, error) {
	_, err := s.influencerRepository.GetByEmail(ctx, input.Email)
	if err == nil {
		return "", influencer.ErrorEmailConflict
	} else if !errors.Is(err, influencer.ErrorNotFound) {
		return "", err
	}

	hashedPassword, err := password.Generate(input.Password)
	if err != nil {
		return "", err
	}

	newInfluencer := influencer.RegisterRequest{
		Name:           input.Name,
		Email:          input.Email,
		Password:       hashedPassword,
		PhoneNumber:    input.PhoneNumber,
		Platforms:      input.Platforms,
		FollowersCount: input.FollowersCount,
		Category:       input.Category,
		Bio:            input.Bio,
		PricePerPost:   input.PricePerPost,
		Address:        input.Address,
	}
	return s.Create(ctx, newInfluencer)
}

func (s *InfluencerService) GetByEmail(ctx context.Context, email string) (res influencer.Response, err error) {
	user, err := s.influencerRepository.GetByEmail(ctx, email)
	if err != nil {
		return
	}
	res = influencer.ParseFromEntity(user)
	return
}

func (s *InfluencerService) GetByEmailEntity(ctx context.Context, email string) (user influencer.Entity, err error) {
	user, err = s.influencerRepository.GetByEmail(ctx, email)
	if err != nil {
		return
	}
	return
}

func (s *InfluencerService) GetByID(ctx context.Context, id string) (res influencer.Response, err error) {
	user, err := s.influencerRepository.GetByID(ctx, id)
	if err != nil {
		return
	}
	res = influencer.ParseFromEntity(user)
	return
}

func (s *InfluencerService) Create(ctx context.Context, data influencer.RegisterRequest) (id string, err error) {
	id, err = s.influencerRepository.Create(ctx, influencer.RegisterRequest{
		Name:           data.Name,
		Email:          data.Email,
		Password:       data.Password,
		PhoneNumber:    data.PhoneNumber,
		Platforms:      data.Platforms,
		FollowersCount: data.FollowersCount,
		Category:       data.Category,
		Bio:            data.Bio,
		PricePerPost:   data.PricePerPost,
		Address:        data.Address,
	})
	if err != nil {
		return
	}

	return
}
