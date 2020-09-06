package web

import (
	"context"
)

type Datastore interface {
	CreateUser(context.Context, string, string, string) (string, error)
	GetUserPassword(context.Context, string) (string, error)
	GetUserInfo(context.Context, string, *User) error
}