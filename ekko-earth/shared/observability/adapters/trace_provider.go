package adapters

import (
	"time"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTraceProvider() (*trace.TracerProvider, error) {
	traceExporter, err := stdouttrace.New(stdouttrace.WithPrettyPrint())

	if err != nil {
		return nil, err
	}

	traceProvider := trace.NewTracerProvider(trace.WithBatcher(traceExporter, trace.WithBatchTimeout(time.Second)))

	return traceProvider, nil
}
