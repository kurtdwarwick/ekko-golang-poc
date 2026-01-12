package adapters

import (
	"go.opentelemetry.io/otel/exporters/stdout/stdoutlog"
	"go.opentelemetry.io/otel/sdk/log"
)

func NewLoggerProvider() (*log.LoggerProvider, error) {
	logExporter, err := stdoutlog.New()

	if err != nil {
		return nil, err
	}

	loggerProvider := log.NewLoggerProvider(log.WithProcessor(log.NewBatchProcessor(logExporter)))

	return loggerProvider, nil
}
