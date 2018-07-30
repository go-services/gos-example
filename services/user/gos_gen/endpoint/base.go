package endpoint

import (
	kitEndpoint "github.com/go-kit/kit/endpoint"
	"context"
	"github.com/go-services/gos-project/services/user/resource"
	"github.com/go-services/core/endpoint"
)

type BaseEndpoints interface {
	Index() endpoint.GOSEndpoint
}

type baseEndpoints struct {
	index endpoint.GOSEndpoint
}

type IndexRequest struct {
	V1 string
}

type IndexResponse struct {
	R1  string `json:"greeting"`
	Err error  `json:"-"`
}

func MakeBaseEndpoints(b resource.Base, global ...kitEndpoint.Middleware) BaseEndpoints {
	global = append(global, globalMiddleware()...)
	endpoints := makeBaseEndpoints(b)
	endpoints.Index().Middleware(global...)
	return endpoints
}

func makeBaseEndpoints(b resource.Base) BaseEndpoints {
	return &baseEndpoints{
		index: makeIndexEndpoint(b),
	}
}
func makeIndexEndpoint(b resource.Base) endpoint.GOSEndpoint {
	return endpoint.NewGOSEndpoint(indexEndpoint(b), indexEndpointMiddleware()...)
}

func indexEndpoint(b resource.Base) kitEndpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(IndexRequest)
		res, err := b.Index(ctx, req.V1)
		if err != nil {
			return nil, err
		}
		return IndexResponse{R1: res}, nil
	}
}

func indexEndpointMiddleware() []kitEndpoint.Middleware {
	return []kitEndpoint.Middleware{}
}

func globalMiddleware() []kitEndpoint.Middleware {
	return []kitEndpoint.Middleware{}
}

func (b *baseEndpoints) Index() endpoint.GOSEndpoint {
	return b.index
}
