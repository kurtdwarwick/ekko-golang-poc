module github.com/ekko-earth/shared/rabbitmq

go 1.25.4

require (
	github.com/ekko-earth/shared/messaging v0.0.0
	github.com/rabbitmq/amqp091-go v1.10.0
)

require github.com/google/uuid v1.6.0 // indirect

replace github.com/ekko-earth/shared/messaging => ../messaging
