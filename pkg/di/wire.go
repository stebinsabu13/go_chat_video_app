//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	http "github.com/stebin13/go_chat_video_app/pkg/api"
	handler "github.com/stebin13/go_chat_video_app/pkg/api/handler"
	config "github.com/stebin13/go_chat_video_app/pkg/config"
	db "github.com/stebin13/go_chat_video_app/pkg/db"
	repository "github.com/stebin13/go_chat_video_app/pkg/repository"
	usecase "github.com/stebin13/go_chat_video_app/pkg/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		repository.NewAuthRepo,
		usecase.NewAuthUsecase,
		handler.NewAuthHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
