package http

import (
	"chat-backend/docs"
	http_handlers "chat-backend/internal/http/handlers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

func (s *Server) useRoutes() {
	apiV1router := s.gin.Group("/api/v1")

	apiV1router.GET("/hello-world", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})

	apiV1router.POST("/chat", http_handlers.AddChat)
}

func (s *Server) useHealth() {
	s.gin.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy")
	})
}

func (s *Server) useSwagger() {
	docs.SwaggerInfo.BasePath = "/"
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
