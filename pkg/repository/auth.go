package repository

import (
	"context"
	"errors"
	"time"

	"github.com/stebin13/go_chat_video_app/pkg/repository/interfaces"
	"github.com/stebin13/go_chat_video_app/pkg/utils"
	"gorm.io/gorm"
)

type authRepo struct {
	DB *gorm.DB
}

func NewAuthRepo(db *gorm.DB) interfaces.AuthRepo {
	return &authRepo{
		DB: db,
	}
}

func (c *authRepo) FindbyEmail(ctx context.Context, email string) (utils.ResponseUsers, error) {
	var user utils.ResponseUsers
	query := `SELECT * from users where email=$1`
	c.DB.Raw(query, email).Scan(&user)
	if user.ID == 0 {
		return user, errors.New("invalid email")
	}
	return user, nil
}

func (cr *authRepo) SignUp(ctx context.Context, body utils.BodySignUpuser) (string, error) {
	var userid string
	query := `insert into users(created_at,updated_at,first_name,last_name,email,mobile_number,password)values($1,$2,$3,$4,$5,$6,$7) returning id`
	if err := cr.DB.Raw(query, time.Now(), time.Now(), body.FirstName, body.LastName, body.Email, body.MobileNum, body.Password).Scan(&userid).Error; err != nil {
		return userid, err
	}
	return userid, nil
}
