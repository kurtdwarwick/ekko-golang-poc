module github.com/ekko-earth/organisation

go 1.25.4

require (
	github.com/ekko-earth/shared/adapters v0.0.0
	github.com/ekko-earth/shared/application v0.0.0
	github.com/ekko-earth/shared/gorm v0.0.0
	github.com/ekko-earth/shared/grpc v0.0.0
	github.com/ekko-earth/shared/http v0.0.0
	github.com/ekko-earth/shared/messaging v0.0.0
	github.com/ekko-earth/shared/observability v0.0.0
	github.com/ekko-earth/shared/outbox v0.0.0
	github.com/ekko-earth/shared/policies v0.0.0
	github.com/ekko-earth/shared/rabbitmq v0.0.0
	github.com/google/uuid v1.6.0
	go.uber.org/mock v0.6.0
	google.golang.org/grpc v1.77.0
	google.golang.org/protobuf v1.36.11
	gorm.io/gorm v1.31.1
)

require (
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	golang.org/x/crypto v0.43.0 // indirect
	golang.org/x/sync v0.17.0 // indirect
)

require (
	filippo.io/edwards25519 v1.1.0 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/denisbrodbeck/machineid v1.0.1 // indirect
	github.com/felixge/httpsnoop v1.0.4 // indirect
	github.com/go-logr/logr v1.4.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-sql-driver/mysql v1.8.1 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	go.opentelemetry.io/auto/sdk v1.2.1 // indirect
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.64.0 // indirect
	go.opentelemetry.io/otel v1.39.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.39.0 // indirect
	go.opentelemetry.io/otel/metric v1.39.0 // indirect
	go.opentelemetry.io/otel/sdk v1.39.0 // indirect
	go.opentelemetry.io/otel/trace v1.39.0 // indirect
	golang.org/x/net v0.46.1-0.20251013234738-63d1a5100f82 // indirect
	golang.org/x/sys v0.39.0 // indirect
	golang.org/x/text v0.30.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20251022142026-3a174f9686a8 // indirect
	gorm.io/datatypes v1.2.7 // indirect
	gorm.io/driver/mysql v1.5.6 // indirect
	gorm.io/driver/postgres v1.6.0 // indirect
)

replace github.com/ekko-earth/shared/policies => ../../shared/policies

replace github.com/ekko-earth/shared/adapters => ../../shared/adapters

replace github.com/ekko-earth/shared/grpc => ../../shared/grpc

replace github.com/ekko-earth/shared/http => ../../shared/http

replace github.com/ekko-earth/shared/gorm => ../../shared/gorm

replace github.com/ekko-earth/shared/rabbitmq => ../../shared/rabbitmq

replace github.com/ekko-earth/shared/messaging => ../../shared/messaging

replace github.com/ekko-earth/shared/application => ../../shared/application

replace github.com/ekko-earth/shared/outbox => ../../shared/outbox

replace github.com/ekko-earth/shared/observability => ../../shared/observability
