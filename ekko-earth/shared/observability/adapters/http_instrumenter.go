package adapters

import (
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/ekko-earth/shared/http/adapters"
)

type HttpInstrumenter struct{}

func (instrumenter *HttpInstrumenter) Instrument(server adapters.HttpServer) http.Handler {
	handler := otelhttp.NewHandler(server.Router, "/")

	server.Server.Handler = handler

	return handler
}
