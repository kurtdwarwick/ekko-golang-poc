package adapters

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/trace"

	"github.com/ekko-earth/shared/http/adapters"
)

type HttpInstrumenter struct{}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{})
}

func newTraceProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(trace.WithBatcher(traceExporter, trace.WithBatchTimeout(time.Second)))

	return traceProvider, nil
}

func (instrumenter *HttpInstrumenter) Instrument(server adapters.HttpServer) http.Handler {
	handler := otelhttp.NewHandler(server.Router, "/")

	server.Server.Handler = handler

	return handler
}

func ConfigureHttpInstrumenter(ctx context.Context) (func(context.Context) error, error) {
	var shutdowns []func(context.Context) error
	var err error

	shutdown := func(ctx context.Context) error {
		var err error

		for _, shutdown := range shutdowns {
			err = errors.Join(err, shutdown(ctx))
		}

		shutdowns = nil

		return err
	}

	handleErr := func(e error) {
		err = errors.Join(e, shutdown(ctx))
	}

	prop := newPropagator()

	otel.SetTextMapPropagator(prop)

	traceProvider, err := newTraceProvider()

	if err != nil {
		handleErr(err)
		return shutdown, err
	}

	shutdowns = append(shutdowns, traceProvider.Shutdown)

	otel.SetTracerProvider(traceProvider)

	return shutdown, err
}
