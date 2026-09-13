package http

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Router struct {
	engine *gin.Engine
}

func NewRouter(logger *zap.Logger) *Router {
	r := gin.Default()
	r.GET("/health", func(c *gin.Context) {
		logger.Info("Health check endpoint hit")
		c.JSON(200, gin.H{
			"status": "healthy",
		})
	})
	// Define your routes here
	return &Router{engine: r}
}

func (r *Router) Run(addr string) error {
	return r.engine.Run(addr)
}
