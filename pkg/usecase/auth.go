package usecase

import (
	"context"
	"errors"
	"log"

	"github.com/stebin13/go_chat_video_app/pkg/repository/interfaces"
	"github.com/stebin13/go_chat_video_app/pkg/support"
	services "github.com/stebin13/go_chat_video_app/pkg/usecase/interfaces"
	"github.com/stebin13/go_chat_video_app/pkg/utils"
)

type authUsecase struct {
	AuthRepo interfaces.AuthRepo
}

func NewAuthUsecase(repo interfaces.AuthRepo) services.AuthUsecase {
	return &authUsecase{
		AuthRepo: repo,
	}
}

func (cr *authUsecase) SignUp(ctx context.Context, body utils.BodySignUpuser) (string, error) {
	if body.FirstName == "" || body.LastName == "" || body.Email == "" || body.MobileNum == "" || body.ConfirmPassword == "" || body.Password == "" {
		return "", errors.New("fill the required fields")
	}
	if body.ConfirmPassword != body.Password {
		return "", errors.New("password mismatch")
	}
	if _, err := cr.AuthRepo.FindbyEmail(ctx, body.Email); err == nil {
		return "", errors.New("user already exsists")
	}
	hashpassword, err1 := support.HashPassword(body.Password)
	if err1 != nil {
		return "", errors.New("error while hashing password")
	}
	body.Password = hashpassword
	log.Println(body)
	userid, err := cr.AuthRepo.SignUp(ctx, body)
	if err != nil {
		return userid, err
	}
	return userid, nil
}
