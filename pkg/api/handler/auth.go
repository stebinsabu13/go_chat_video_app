package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/stebin13/go_chat_video_app/pkg/usecase/interfaces"
	"github.com/stebin13/go_chat_video_app/pkg/utils"
)

type AuthHandler struct {
	AuthUsecase services.AuthUsecase
}

func NewAuthHandler(usecase services.AuthUsecase) *AuthHandler {
	return &AuthHandler{
		AuthUsecase: usecase,
	}
}

// func (cr *AuthHandler) LoginHandler(c *gin.Context) {
// 	body := utils.BodyLogin{
// 		Email:    c.PostForm("email"),
// 		Password: c.PostForm("password"),
// 	}
// }

func (cr *AuthHandler) SignUp(c *gin.Context) {
	body := utils.BodySignUpuser{
		FirstName:       c.PostForm("fName"),
		LastName:        c.PostForm("lName"),
		Email:           c.PostForm("email"),
		MobileNum:       c.PostForm("mobile_num"),
		Password:        c.PostForm("Password"),
		ConfirmPassword: c.PostForm("confirmPassword"),
	}
	userid, err := cr.AuthUsecase.SignUp(c.Request.Context(), body)
	if err != nil {
		c.HTML(http.StatusBadRequest, "signup.html", err)
		return
	}
	c.HTML(http.StatusOK, "signup.html", userid)
}

func (cr *AuthHandler) LoginPage(c *gin.Context) {
	c.HTML(200, "login.html", "Please login")
}

func (cr *AuthHandler) SignUpPage(c *gin.Context) {
	c.HTML(http.StatusOK, "signup.html", "Please login")
}
