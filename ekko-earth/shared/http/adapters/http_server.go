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

func (server *HttpServer) Start(context context.Context) error {
	slog.Info("Starting HTTP consumer")

	go func() {
		if err := server.Server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			slog.Error("Failed to listen and serve", "error", err)
		}
	}()

	return nil
}

func (server *HttpServer) Stop(context context.Context) error {
	slog.Info("Stopping HTTP consumer")

	err := server.Server.Shutdown(context)

	return err
}
