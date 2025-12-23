package adapters

import "context"

type DatabaseConfiguration struct {
	Host     string
	Port     int
	Username string
	Password string
	Database string
	Schema   string
}

type Database interface {
	Connect(ctx context.Context) error
	Disconnect(ctx context.Context) error
}
