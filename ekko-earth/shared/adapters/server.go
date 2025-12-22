package adapters

import "context"

type Server interface {
	Start(context context.Context) error
	Stop(context context.Context) error
}
