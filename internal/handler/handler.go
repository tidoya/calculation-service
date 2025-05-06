package handler

import (
	"calculation-service/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// @FIX-ME
// GIN Есть вопрос стоит ли использовать фреймворк gin для работы с http или пользоваться нативом
func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	// AUTH routes
	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
		auth.POST("/logout", h.logout)
	}

	return router

}
