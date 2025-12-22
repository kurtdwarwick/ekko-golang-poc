package adapters

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/gorilla/mux"
)

type HttpServer struct {
	Server *http.Server
	Router *mux.Router
}

type HttpServerConfiguration struct {
	Address string
}

func NewHttpServer(configuration HttpServerConfiguration) *HttpServer {
	router := mux.NewRouter().StrictSlash(true)

	httpServer := &http.Server{
		Addr:    configuration.Address,
		Handler: router,
	}

	return &HttpServer{Server: httpServer, Router: router}
}

func (server *HttpServer) Start() {
	slog.Info("Starting HTTP consumer")

	go func() {
		if err := server.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to listen and serve", "error", err)
		}
	}()
}

func (server *HttpServer) Stop() {
	slog.Info("Stopping HTTP consumer")

	server.Server.Shutdown(context.TODO())
}
