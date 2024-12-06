package interfaces

import (
	"Marketplace/internal/domain/ad"
	"context"
)

// AdService определяет интерфейс для работы с объявлениями.
type AdService interface {
	// Create создает новое объявление.
	Create(ctx context.Context, input ad.CreateRequest) (string, error)

	// GetByID возвращает объявление по его ID.
	GetByID(ctx context.Context, id string) (ad.Response, error)

	// GetAll возвращает все объявления.
	GetAll(ctx context.Context) ([]ad.Response, error)

	// UpdateStatus обновляет статус объявления.
	UpdateStatus(ctx context.Context, id string, status string) error
}
