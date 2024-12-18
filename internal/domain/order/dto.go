package order

import (
	"time"
)

type CreateRequest struct {
	AdID         string  `json:"ad_id" binding:"required"`
	CompanyID    string  `json:"company_id" binding:"required"`
	InfluencerID string  `json:"influencer_id" binding:"required"`
	Price        float64 `json:"price" binding:"required"`
	Description  string  `json:"description"`
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
	Price        float64   `json:"price"`
	Description  string    `json:"description,omitempty"`
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
		Price:        entity.Price,
		Description:  entity.Description,
		CreatedAt:    entity.CreatedAt,
		UpdatedAt:    entity.UpdatedAt,
	}
}
