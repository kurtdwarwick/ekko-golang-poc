package main

import (
	"context"
	"fmt"
)

type Consumer interface {
	Consume(message any, ctx context.Context) error
}

type HttpConsumer struct{}

func NewHttpConsumer(route string, methods []string) *HttpConsumer {
	return &HttpConsumer{}
}

func (consumer *HttpConsumer) Consume(message any, ctx context.Context) error {
	fmt.Println("Consuming message")
	return nil
}

func (consumer *HttpConsumer) Handle(message any, ctx context.Context) error {
	consumer.Consume(message, ctx)
	return nil
}

type CustomHttpConsumer struct {
	HttpConsumer
}

func NewCustomHttpConsumer(route string, methods []string) *CustomHttpConsumer {
	return &CustomHttpConsumer{HttpConsumer: *NewHttpConsumer(route, methods)}
}

func (consumer *CustomHttpConsumer) Consume(message any, ctx context.Context) error {
	fmt.Println("Consuming custom message")
	return nil
}

func main() {
	consumer := NewCustomHttpConsumer("/test", []string{"GET"})
	consumer.Handle(nil, context.Background())
}
