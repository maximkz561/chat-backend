package http

import (
	"chat-backend/core"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	gin *gin.Engine
}

func NewServer(g *gin.Engine) *Server {
	server := Server{
		gin: g,
	}

	server.useHealth()
	server.useRoutes()
	server.useSwagger()

	return &server
}

func (s *Server) Start() error {
	return s.gin.Run(fmt.Sprintf(":%s", core.Settings.HttpPort))
}

func (s *Server) Shutdown(ctx context.Context) error {
	// For gin, there's no built-in graceful shutdown, we'd need to use the http package to do this.
	// See https://github.com/gin-gonic/gin/issues/296
	return nil
}

func (s *Server) Fatal(err error) {
	// Gin does not have a logger by default, you need to set it up yourself if you want logging
	// s.gin.Logger.Fatal(err)
}
