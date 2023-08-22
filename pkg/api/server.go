package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stebin13/go_chat_video_app/pkg/api/handler"
	"github.com/stebin13/go_chat_video_app/pkg/api/routes"
	"github.com/stebin13/go_chat_video_app/pkg/config"
)

type ServerHTTP struct {
	engine *gin.Engine
}

func NewServerHTTP(authHandler *handler.AuthHandler) *ServerHTTP {

	engine := gin.New()

	engine.Use(gin.Logger())

	// swagger docs
	// engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// set up routes
	routes.UserRoutes(engine.Group("/"), authHandler)
	// routes.AdminRoutes(engine.Group("/"), adminHandler, productHandler, orderHandler)

	// no handler
	engine.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"StatusCode": 404,
			"msg":        "invalid url",
		})
	})

	return &ServerHTTP{engine: engine}
}

func (s *ServerHTTP) Start() {
	port := config.GetLocalPort()
	s.engine.LoadHTMLGlob("static/*.html")
	if err := s.engine.Run(":" + port); err != nil {
		return
	}
}
