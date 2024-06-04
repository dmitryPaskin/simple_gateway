package router

import (
	"API/gateway/internal/controller"
	"context"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Router struct {
	chi      *chi.Mux
	handlers controller.Controller
}

func New() Router {
	return Router{
		chi:      chi.NewRouter(),
		handlers: controller.Controller{},
	}
}

func (r *Router) Start() {
	r.chi.Get("/service1", r.handlers.GetMessageService1)
	r.chi.Get("/service2", r.handlers.GetMessageService2)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      r.chi,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	sigChan := make(chan os.Signal, 1)
	defer close(sigChan)
	signal.Notify(sigChan, syscall.SIGINT)

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server error: %v", err)
		}
	}()

	<-sigChan
	stopCTX, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	if err := server.Shutdown(stopCTX); err != nil {
		log.Fatalf("Server shutdown error: %v", err)
	}
}
