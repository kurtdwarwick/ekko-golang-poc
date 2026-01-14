package http

import (
	"github.com/ekko-earth/shared/http/adapters"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
)

func Instrument(server adapters.HttpServer) {
	handler := otelhttp.NewHandler(server.Router, "/")

	server.Server.Handler = handler
}
