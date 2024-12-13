package ad

import (
	"errors"
	"time"
)

var (
	ErrorNotFound     = errors.New("ad not found")
	ErrorInvalidTitle = errors.New("title is invalid or empty")
	ErrorInvalidPrice = errors.New("price must be greater than zero")
)

type CreateRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func ParseFromEntity(entity Entity) Response {
	return Response{
		ID:          entity.ID,
		Title:       entity.Title,
		Description: entity.Description,
		Price:       entity.Price,
		Status:      entity.Status,
		CreatedAt:   entity.CreatedAt,
		UpdatedAt:   entity.UpdatedAt,
	}
}

func ParseFromEntities(entities []Entity) []Response {
	var responses []Response
	for _, entity := range entities {
		responses = append(responses, ParseFromEntity(entity))
	}
	return responses
}

func (r *CreateRequest) Validate() error {
	if r.Title == "" {
		return ErrorInvalidTitle
	}
	if r.Price <= 0 {
		return ErrorInvalidPrice
	}
	return nil
}
