package interfaces

import (
	"context"

	"github.com/stebin13/go_chat_video_app/pkg/utils"
)

type AuthRepo interface {
	SignUp(context.Context, utils.BodySignUpuser) (string, error)
	FindbyEmail(context.Context, string) (utils.ResponseUsers, error)
}
