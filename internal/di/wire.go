//go:build wireinject
// +build wireinject

package di

import (
	http "Marketplace/internal/api"
	"Marketplace/internal/api/handler"
	"Marketplace/internal/config"
	"Marketplace/internal/db"
	"Marketplace/internal/repository"
	"Marketplace/internal/service"
	"Marketplace/internal/utils/jwt"
	"github.com/google/wire"
	_ "github.com/lib/pq"
)

func InitializeAPI(cfg config.Config) (*http.Server, error) {
	wire.Build(
		db.ConnectDatabase,
		repository.NewCompanyRepository,
		service.NewCompanyService,
		handler.NewCompanyHandler,
		jwt.ProvideSecretKey,
		http.NewServer,
	)
	return &http.Server{}, nil
}
