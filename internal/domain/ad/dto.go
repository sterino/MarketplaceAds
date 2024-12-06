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

// Структура для создания нового объявления
type CreateRequest struct {
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// Структура ответа на запрос
type Response struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Price       float64   `json:"price"`
	Status      string    `json:"status"`
	OrdersID    []string  `json:"orders"`
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
		OrdersID:    entity.OrdersID,
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

// Валидация для создания нового объявления
func (r *CreateRequest) Validate() error {
	if r.Title == "" {
		return ErrorInvalidTitle
	}
	if r.Price <= 0 {
		return ErrorInvalidPrice
	}
	return nil
}
