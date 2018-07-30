package resource

import (
	"context"
)

// @resource(name="base", route="/")
// @endpoint_middleware(name="github.com/sdasd/asd.MyMiddleware")
type Base interface {
	// @http(method="GET", route="/")
	// @http(method="GET", route="/{username}")
	// @http(method="GET", route="/{username}/")
	// @param(name="username", in="url")
	Index(ctx context.Context, username string) (greeting string, err error)
}

// @stub
type base struct {
}

func NewBase() Base {
	return &base{}
}

func (base) Index(ctx context.Context, username string) (greeting string, err error) {
	greeting = "Hello " + username
	return
}
