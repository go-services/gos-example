package gos_gen

import (
	"github.com/go-services/gos-project/services/user/gos_gen/resource"
	"github.com/go-services/core"
	"github.com/kevinburke/go.uuid"
)

func NewServiceConfig() core.ServiceConfig {
	return core.ServiceConfig{
		Name:        "user",
		HTTPAddress: ":8080",
		GRPCAddress: ":8081",
		Instance:    uuid.NewV4().String(),
	}
}

type UserService interface {
	Base() *resource.Base
}
