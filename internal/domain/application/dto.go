package application

import (
	"errors"
	"time"
)

var (
	ErrorNotFound = errors.New("ad not found")
)

type CreateRequest struct {
	AdID         string `json:"ad_id" binding:"required"`
	CompanyID    string `json:"company_id" binding:"required"`
	InfluencerID string `json:"influencer_id" binding:"required"`
}

type UpdateStatusRequest struct {
	Status string `json:"status" binding:"required"`
}

type Response struct {
	ID           string    `json:"id"`
	AdID         string    `json:"ad_id"`
	CompanyID    string    `json:"company_id"`
	InfluencerID string    `json:"influencer_id"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:           entity.ID,
		AdID:         entity.AdID,
		CompanyID:    entity.CompanyID,
		InfluencerID: entity.InfluencerID,
		Status:       entity.Status,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}

func ParseEntities(entities []Entity) []Response {
	responses := make([]Response, len(entities))
	for i, entity := range entities {
		responses[i] = ParseFromEntity(entity)
	}
	return responses
}
