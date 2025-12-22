module github.com/ekko-earth/impact

go 1.25.4

require (
	github.com/ekko-earth/shared/adapters v0.0.0
	github.com/ekko-earth/shared/gorm v0.0.0-00010101000000-000000000000
	github.com/ekko-earth/shared/messaging v0.0.0
	github.com/ekko-earth/shared/rabbitmq v0.0.0
	github.com/google/uuid v1.6.0
	gorm.io/gorm v1.31.1
)

require (
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20240606120523-5a60cdf6a761 // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.2 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/rabbitmq/amqp091-go v1.10.0 // indirect
	golang.org/x/crypto v0.31.0 // indirect
	golang.org/x/sync v0.10.0 // indirect
	golang.org/x/text v0.21.0 // indirect
	gorm.io/driver/postgres v1.6.0 // indirect
)

replace github.com/ekko-earth/shared/policies => ../../shared/policies

replace github.com/ekko-earth/shared/adapters => ../../shared/adapters

replace github.com/ekko-earth/shared/grpc => ../../shared/grpc

replace github.com/ekko-earth/shared/http => ../../shared/http

replace github.com/ekko-earth/shared/gorm => ../../shared/gorm

replace github.com/ekko-earth/shared/rabbitmq => ../../shared/rabbitmq

replace github.com/ekko-earth/shared/messaging => ../../shared/messaging
