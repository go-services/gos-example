package user

import (
	gosResource "github.com/go-services/gos-project/services/user/gos_gen/resource"
	gosEndpoint "github.com/go-services/gos-project/services/user/gos_gen/endpoint"

	"github.com/go-services/gos-project/services/user/gos_gen"
	"github.com/go-kit/kit/log"
	"github.com/go-services/gos-project/services/user/resource"
)

// this is the service struct, it must implement the generated interface gos_gen.UserService
// we create this structure so we can have a way to create and modify all of the resources of the service
// @service(name="service", route="/")
type user struct {
	logger log.Logger
}

// @new()
func NewUser(logger log.Logger) gos_gen.UserService {
	return &user{
		logger: logger,
	}
}

func (u *user) Base() *gosResource.Base {
	// create resource instance
	// you can add resource middleware as many as you like here or using annotations
	b := gosResource.NewBase(resource.NewBase())

	// create endpoints
	// you can add endpoint middleware as many as you like here or using annotations
	endpoints := gosEndpoint.MakeBaseEndpoints(b)
	// here you can plug in your endpoint function if you do not want to use the default, be aware that if you do that
	// you most likely will have to also update the transport decoder
	//  e.x endpoints.Index().SetEndpoint(EndpointFunc)
	// you can also add endpoint middleware to specific endpoints if you want to.
	//  e.x endpoints.Index().Middleware(SomeMiddleware)

	// create resource transports
	// you can add transport options here if needed
	base := gosResource.MakeBaseTransports(endpoints, nil)
	// here you can modify specific transports for methods,
	// like change the handler, change the encoder, change the decoder
	// e.x base.Http().Index().SetDecoder(DecoderFunc)
	// e.x base.Http().Index().SetEncoder(EncoderFunc)
	// e.x base.Http().Index().SetHandle(Handler)
	// e.x base.Http().Index().Options(ServerOptions)
	return base
}
