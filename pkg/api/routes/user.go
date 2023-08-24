package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/stebin13/go_chat_video_app/pkg/api/handler"
)

func UserRoutes(api *gin.RouterGroup, authHandler *handler.AuthHandler) {
	login := api.Group("/user")
	{
		// Request JWT
		login.GET("/login", authHandler.LoginPage)
		login.POST("/login", authHandler.LoginHandler)
	}
	signup := api.Group("/user")
	{
		signup.GET("/signup", authHandler.SignUpPage)
		signup.POST("/signup", authHandler.SignUp)
		// signup.POST("/signup/otp/verify", authHandler.SignupOtpverify)
	}
}
