package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/rpuggian/ports-service/api/domain/contracts"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
	wait       time.Duration
}

func NewMuxRouter(portHandler contracts.PortHandler) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/ports/{id}", portHandler.GetFileByID).Methods(http.MethodGet)
	r.HandleFunc("/ports", portHandler.UploadPortFileHandler).Methods(http.MethodPost)
	return r
}

func NewServer(apiPort string, portHandler contracts.PortHandler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr: fmt.Sprintf(":%s", apiPort),
			// Good practice to set timeouts to avoid Slowloris attacks.
			WriteTimeout: time.Second * 15,
			ReadTimeout:  time.Second * 15,
			Handler:      NewMuxRouter(portHandler),
		},
		wait: 10 * time.Second,
	}
}

func (s *Server) Serve() {
	go func() {
		err := s.httpServer.ListenAndServe()
		if err != nil {
			log.Printf("Listen and serve: %v", err)
		}
	}()
}

func (s *Server) GracefulShutdown() {
	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), s.wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	err := s.httpServer.Shutdown(ctx)
	if err != nil {
		log.Printf("error to shutdown server: %v\n", err)
	}
}
