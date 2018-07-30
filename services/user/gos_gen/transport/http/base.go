package http

import (
	"github.com/go-services/gos-project/services/user/gos_gen/endpoint"
	goKitEndpoint "github.com/go-kit/kit/endpoint"
	goHttp "net/http"
	"context"
	"encoding/json"
	"github.com/go-services/core/transport/http"
	goKitHttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

type BaseHttp interface {
	Index() http.Http
	Iterate() []http.Http
}

type baseHttp struct {
	methods map[string]http.Http
}

func MakeBaseHttpTransport(endpoints endpoint.BaseEndpoints, global ...goKitHttp.ServerOption) BaseHttp {
	global = append(global, globalServerOptions()...)
	httpTransport := makeBaseHttpTransport(endpoints)
	httpTransport.Index().Options(global)
	return httpTransport
}
func makeBaseHttpTransport(endpoints endpoint.BaseEndpoints) BaseHttp {
	return &baseHttp{
		methods: map[string]http.Http{
			"index": makeIndexHttpTransport(endpoints.Index().Endpoint()),
		},
	}
}

func makeIndexHttpTransport(endpoint goKitEndpoint.Endpoint) http.Http {
	return http.NewHttpTransport([]http.MethodRoute{
		{
			Name:   "base:index:no_name",
			Route:  "/",
			Method: http.GET,
		}, {
			Name:   "base:index",
			Method: http.GET,
			Route:  "/{username}",
		}, {
			Name:   "base:index:no_name",
			Route:  "/{username}/",
			Method: http.GET,
		},
	}, indexDecoder, indexEncoder, endpoint, indexOptions()...)
}

func indexDecoder(ctx context.Context, r *goHttp.Request) (request interface{}, err error) {
	req := endpoint.IndexRequest{}
	req.V1 = mux.Vars(r)["username"]
	return req, err
}

func indexEncoder(ctx context.Context, w goHttp.ResponseWriter, response interface{}) error {
	resp := response.(endpoint.IndexResponse)
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO handle error
	return json.NewEncoder(w).Encode(resp.R1)
}

func indexOptions() []goKitHttp.ServerOption {
	return []goKitHttp.ServerOption{}
}

func (b *baseHttp) Index() http.Http {
	return b.methods["index"]
}

func (b *baseHttp) Iterate() (methods []http.Http) {
	for _, v := range b.methods {
		methods = append(methods, v)
	}
	return
}

func globalServerOptions() []goKitHttp.ServerOption {
	return []goKitHttp.ServerOption{}
}
