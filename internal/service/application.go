package service

import (
	"Marketplace/internal/domain/application"
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type ApplicationService struct {
	applicationRepo interfaces.ApplicationRepository
}

func NewApplicationService(repo interfaces.ApplicationRepository) services.ApplicationService {
	return &ApplicationService{applicationRepo: repo}
}

func (s *ApplicationService) Create(ctx context.Context, req application.CreateRequest) (application.Response, error) {
	id := uuid.New().String()
	entity := application.Entity{
		ID:           id,
		AdID:         req.AdID,
		CompanyID:    req.CompanyID,
		InfluencerID: req.InfluencerID,
		Status:       "pending",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdID, err := s.applicationRepo.Create(ctx, entity)
	if err != nil {
		return application.Response{}, err
	}

	entity.ID = createdID
	return application.ParseFromEntity(entity), nil
}

func (s *ApplicationService) GetByID(ctx context.Context, id string) (application.Response, error) {
	entity, err := s.applicationRepo.GetByID(ctx, id)
	if err != nil {
		return application.Response{}, err
	}
	return application.ParseFromEntity(entity), nil
}

func (s *ApplicationService) GetByAdID(ctx context.Context, adID string) ([]application.Response, error) {
	entities, err := s.applicationRepo.GetByAdID(ctx, adID)
	if err != nil {
		return nil, err
	}
	return application.ParseEntities(entities), nil
}

func (s *ApplicationService) GetByInfluencerID(ctx context.Context, influencerID string) ([]application.Response, error) {
	entities, err := s.applicationRepo.GetByInfluencerID(ctx, influencerID)
	if err != nil {
		return nil, err
	}
	return application.ParseEntities(entities), nil
}

func (s *ApplicationService) UpdateStatus(ctx context.Context, id, status string) error {
	validStatuses := map[string]bool{"pending": true, "approved": true, "rejected": true}
	if !validStatuses[status] {
		return errors.New("invalid status")
	}
	return s.applicationRepo.UpdateStatus(ctx, id, status)
}

func (s *ApplicationService) Delete(ctx context.Context, id string) error {
	return s.applicationRepo.Delete(ctx, id)
}
