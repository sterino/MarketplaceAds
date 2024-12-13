// Code generated by Wire. DO NOT EDIT.

//go:build !wireinject
// +build !wireinject

package di

import (
	"Marketplace/internal/api"
	"Marketplace/internal/api/handler"
	"Marketplace/internal/config"
	"Marketplace/internal/db"
	"Marketplace/internal/repository"
	"Marketplace/internal/service"
	"Marketplace/internal/utils/jwt"
	_ "github.com/lib/pq"
)

func InitializeAPI(cfg config.Config) (*api.Server, error) {

	sqlxDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}

	err = db.Migrate(sqlxDB)
	if err != nil {
		return nil, err
	}

	applicationRepository := repository.NewApplicationRepository(sqlxDB)
	companyRepository := repository.NewCompanyRepository(sqlxDB)
	orderRepository := repository.NewOrderRepository(sqlxDB)
	influencerRepository := repository.NewInfluencerRepository(sqlxDB)
	codeRepository := repository.NewCodeRepository(sqlxDB)
	userRepository := repository.NewUserRepository(sqlxDB)
	adRepository := repository.NewAdRepository(sqlxDB)

	secretKey := jwt.ProvideSecretKey()

	applicationService := service.NewApplicationService(applicationRepository)
	orderService := service.NewOrderService(orderRepository)
	companyService := service.NewCompanyService(companyRepository, codeRepository, secretKey)
	influencerService := service.NewInfluencerService(influencerRepository, secretKey, codeRepository)
	userService := service.NewUserService(userRepository)
	adService := service.NewAdService(adRepository)

	applicationHandler := handler.NewApplicationHandler(applicationService)
	orderHandler := handler.NewOrderHandler(orderService)
	companyHandler := handler.NewCompanyHandler(companyService)
	influencerHandler := handler.NewInfluencerHandler(influencerService)
	userHandler := handler.NewUserHandler(userService)
	adHandler := handler.NewAdHandler(adService)

	server := api.NewServer(companyHandler, influencerHandler, userHandler, adHandler, orderHandler, applicationHandler)
	return server, nil
}
