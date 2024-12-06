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
	// Здесь можно добавить дополнительную валидацию или обработку данных
	id, err := s.adRepository.Create(ctx, input)
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
	// Преобразуем сущность в response
	res := ad.ParseFromEntity(adEntity)
	return res, nil
}

func (s *AdService) GetAll(ctx context.Context) ([]ad.Response, error) {
	ads, err := s.adRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	// Преобразуем все сущности в response
	res := ad.ParseFromEntities(ads)
	return res, nil
}

func (s *AdService) UpdateStatus(ctx context.Context, id string, status string) error {
	// Проверяем, существует ли объявление перед обновлением
	_, err := s.adRepository.GetByID(ctx, id)
	if err != nil {
		return err
	}

	// Обновляем статус объявления
	return s.adRepository.UpdateStatus(ctx, id, status)
}
