package resource

import (
	"github.com/go-services/gos-project/services/user/gos_gen/endpoint"
	"github.com/go-services/gos-project/services/user/gos_gen/transport/http"
	"github.com/go-services/gos-project/services/user/resource"
	goKitHttp "github.com/go-kit/kit/transport/http"
)

type Base struct {
	http http.BaseHttp
}

// TODO add GRPC here
func MakeBaseTransports(endpoints endpoint.BaseEndpoints, httpOptions []goKitHttp.ServerOption) *Base {
	httpTransport := http.MakeBaseHttpTransport(endpoints, httpOptions...)
	return &Base{
		http: httpTransport,
	}
}

func (b *Base) Http() http.BaseHttp {
	return b.http
}

type Middleware func(resource.Base) resource.Base

func NewBase(b resource.Base, mdw ...Middleware) resource.Base {
	mdw = append(mdw, defaultBaseMiddleware()...)
	for _, m := range mdw {
		b = m(b)
	}
	return b
}

func defaultBaseMiddleware() []Middleware {
	return []Middleware{}
}
