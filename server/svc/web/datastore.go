package web

import (
	"context"
)

type Datastore interface {
	CreateUser(context.Context, string, string) error
}