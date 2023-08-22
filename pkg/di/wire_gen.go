// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package di

import (
	"github.com/stebin13/go_chat_video_app/pkg/api"
	"github.com/stebin13/go_chat_video_app/pkg/api/handler"
	"github.com/stebin13/go_chat_video_app/pkg/config"
	"github.com/stebin13/go_chat_video_app/pkg/db"
	"github.com/stebin13/go_chat_video_app/pkg/repository"
	"github.com/stebin13/go_chat_video_app/pkg/usecase"
)

// Injectors from wire.go:

func InitializeAPI(cfg config.Config) (*api.ServerHTTP, error) {
	gormDB, err := db.ConnectDatabase(cfg)
	if err != nil {
		return nil, err
	}
	authRepo := repository.NewAuthRepo(gormDB)
	authUsecase := usecase.NewAuthUsecase(authRepo)
	authHandler := handler.NewAuthHandler(authUsecase)
	serverHTTP := api.NewServerHTTP(authHandler)
	return serverHTTP, nil
}