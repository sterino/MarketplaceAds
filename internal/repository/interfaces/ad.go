package interfaces

import (
	"Marketplace/internal/domain/ad"
	"context"
)

// AdRepository интерфейс для работы с репозиторием объявлений
type AdRepository interface {
	// Create создает новое объявление в базе данных
	Create(ctx context.Context, data ad.CreateRequest) (string, error)

	// GetByID возвращает объявление по его ID
	GetByID(ctx context.Context, id string) (ad.Entity, error)

	// GetAll возвращает все объявления
	GetAll(ctx context.Context) ([]ad.Entity, error)

	// UpdateStatus обновляет статус объявления
	UpdateStatus(ctx context.Context, id string, status string) error
}
