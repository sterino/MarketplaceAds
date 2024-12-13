package service

import (
	"Marketplace/internal/domain/order"
	"Marketplace/internal/repository/interfaces"
	services "Marketplace/internal/service/interfaces"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type OrderService struct {
	orderRepo interfaces.OrderRepository
}

func NewOrderService(orderRepo interfaces.OrderRepository) services.OrderService {
	return &OrderService{
		orderRepo: orderRepo,
	}
}

func (s *OrderService) Create(ctx context.Context, req order.CreateRequest) (order.Response, error) {
	id := uuid.New().String()
	entity := order.Entity{
		ID:           id,
		AdID:         req.AdID,
		CompanyID:    req.CompanyID,
		InfluencerID: req.InfluencerID,
		Status:       "pending",
		Price:        req.Price,
		Description:  req.Description,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	createdID, err := s.orderRepo.Create(ctx, entity)
	if err != nil {
		return order.Response{}, err
	}

	entity.ID = createdID
	return order.ParseFromEntity(entity), nil
}

func (s *OrderService) GetByID(ctx context.Context, id string) (order.Response, error) {
	entity, err := s.orderRepo.GetByID(ctx, id)
	if err != nil {
		return order.Response{}, err
	}
	return order.ParseFromEntity(entity), nil
}

func (s *OrderService) UpdateStatus(ctx context.Context, id string, status string) error {
	validStatuses := map[string]bool{"pending": true, "approved": true, "rejected": true, "completed": true}
	if !validStatuses[status] {
		return errors.New("invalid status")
	}

	return s.orderRepo.UpdateStatus(ctx, id, status)
}

func (s *OrderService) GetByCompanyID(ctx context.Context, companyID string) ([]order.Response, error) {
	entities, err := s.orderRepo.GetByCompanyID(ctx, companyID)
	if err != nil {
		return nil, err
	}

	responses := make([]order.Response, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, order.ParseFromEntity(entity))
	}

	return responses, nil
}

func (s *OrderService) GetByInfluencerID(ctx context.Context, influencerID string) ([]order.Response, error) {
	entities, err := s.orderRepo.GetByInfluencerID(ctx, influencerID)
	if err != nil {
		return nil, err
	}

	responses := make([]order.Response, 0, len(entities))
	for _, entity := range entities {
		responses = append(responses, order.ParseFromEntity(entity))
	}

	return responses, nil
}

func (s *OrderService) Delete(ctx context.Context, id string) error {
	return s.orderRepo.Delete(ctx, id)
}
