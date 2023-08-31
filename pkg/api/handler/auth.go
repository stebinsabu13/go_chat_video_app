package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go_chat_video_app/pkg/auth"
	"github.com/stebin13/go_chat_video_app/pkg/support"
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

func (cr *AuthHandler) LoginHandler(c *gin.Context) {

	body := utils.BodyLogin{
		Email:    c.PostForm("email"),
		Password: c.PostForm("password"),
	}
	user, err := cr.AuthUsecase.FindbyEmail(c.Request.Context(), body.Email)
	if err != nil {
		c.HTML(200, "login.html", err)
		return
	}
	ok := support.CheckPasswordHash(body.Password, user.Password)
	if !ok {
		c.HTML(200, "login.html", "Wrong password")
		return
	}

	//generating jwt token for authorization
	tokenString, err1 := auth.GenerateJWT(user.ID)
	if err1 != nil {
		c.HTML(200, "login.html", err)
		return
	}

	//setting cookie for a defined time frame
	c.SetCookie("user-token", tokenString, int(time.Now().Add(5*time.Minute).Unix()), "/", "localhost", false, true)
	c.Redirect(http.StatusFound, "/user/home")
}

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

func (cr *AuthHandler) HomePage(c *gin.Context) {
	c.HTML(200, "home.html", nil)
}

func (cr *AuthHandler) SearchResults(c *gin.Context) {

}
