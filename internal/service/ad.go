package service

import (
	"Marketplace/internal/domain/ad"
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"context"
	"errors"
)

type AdService struct {
	adRepository interfaces.AdRepository
}

func NewAdService(adRepository interfaces.AdRepository) services.AdService {
	return &AdService{
		adRepository: adRepository,
	}
}

func (s *AdService) Create(ctx context.Context, input ad.CreateRequest) (string, error) {
	id, err := s.adRepository.Create(ctx, ad.CreateRequest{
		CompanyID:   input.CompanyID,
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (s *AdService) GetByID(ctx context.Context, id string) (ad.Response, error) {
	adEntity, err := s.adRepository.GetByID(ctx, id)
	if err != nil {
		if errors.Is(err, ad.ErrorNotFound) {
			return ad.Response{}, err
		}
		return ad.Response{}, err
	}

	res := ad.ParseFromEntity(adEntity)
	return res, nil
}

func (s *AdService) GetAll(ctx context.Context) ([]ad.Response, error) {
	ads, err := s.adRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	res := ad.ParseFromEntities(ads)
	return res, nil
}

func (s *AdService) UpdateStatus(ctx context.Context, id string, status string) error {
	_, err := s.adRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	return s.adRepository.UpdateStatus(ctx, id, status)
}

func (s *AdService) GetByCompanyID(ctx context.Context, companyID string) ([]ad.Response, error) {
	ads, err := s.adRepository.GetByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}
	return ad.ParseFromEntities(ads), nil
}

func (s *AdService) Delete(ctx context.Context, id string) error {
	return s.adRepository.Delete(ctx, id)
}
