package observability

import (
	"context"
	"errors"

	"github.com/ekko-earth/shared/observability/adapters"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var Tracer trace.Tracer
var Instrument bool = false

func NewInstrumentation(ctx context.Context, serviceName string) (func(context.Context) error, error) {
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

	prop := adapters.NewPropagator()

	otel.SetTextMapPropagator(prop)

	traceProvider, err := adapters.NewTraceProvider(ctx, serviceName)

	if err != nil {
		handleErr(err)
		return shutdown, err
	}

	shutdowns = append(shutdowns, traceProvider.Shutdown)

	otel.SetTracerProvider(traceProvider)

	Tracer = traceProvider.Tracer(serviceName)
	Instrument = true

	return shutdown, err
}
