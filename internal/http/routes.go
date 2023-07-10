package http

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (s *Server) useRoutes() {
	apiV1router := s.gin.Group("/api/v1")

	apiV1router.GET("/hello-world", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world!")
	})
}

func (s *Server) useHealth() {
	s.gin.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "Healthy")
	})
}
