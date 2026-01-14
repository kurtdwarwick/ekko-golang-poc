package observability

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
)

func PropagateContext(ctx context.Context, headers map[string]any) context.Context {
	carrier := propagation.MapCarrier{}

	for key, value := range headers {
		if value, ok := value.(string); ok {
			carrier.Set(key, value)
		}
	}

	return otel.GetTextMapPropagator().Extract(ctx, carrier)
}

func ExtractFromContext(
	ctx context.Context,
) map[string]any {
	carrier := propagation.MapCarrier{}

	propagator := otel.GetTextMapPropagator()
	propagator.Inject(ctx, carrier)

	keys := carrier.Keys()

	headers := make(map[string]any)

	for _, key := range keys {
		value := carrier.Get(key)
		headers[key] = value
	}

	return headers
}
