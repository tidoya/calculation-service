package handler

import (
	"calculation-service/service"
	"net/http"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

// @FIX-ME
// GIN Есть вопрос стоит ли использовать фреймворк gin для работы с http или пользоваться нативом
func (h *Handler) InitRoutes() *http.ServeMux {
	router := http.NewServeMux()

	// AUTH routes
	router.HandleFunc("/auth/sign-up", h.signUp)
	router.HandleFunc("/auth/sign-in", h.signIn)
	router.HandleFunc("/auth/logout", h.logout)

	// router.HandleFunc("/auth/logout", nil)

	return router

}
