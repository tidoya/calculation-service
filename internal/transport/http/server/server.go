package Calculation

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) RunServer(port string, handler http.Handler) error {
	fmt.Printf("Сервер запущен на http://localhost%s\n", ":"+port)

	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) ShutDownServer(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
