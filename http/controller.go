package http

import (
	"github.com/danielgtaylor/huma/v2"
	"github.com/samber/lo"
	"go.uber.org/fx"
)

type Controller interface {
	RegisterRoutes(app huma.API)
}

func AsController[T any](endpoint T) interface{} {
	return fx.Annotate(
		endpoint,
		fx.ResultTags(`group:"controller"`),
		fx.As(new(Controller)),
	)
}

type RegisterEndpointParameter struct {
	fx.In
	Endpoint []Controller `group:"endpoint"`
	Openapi  huma.API
}

func registerEndpoint(parameters RegisterEndpointParameter) {
	endpoints, openapi := parameters.Endpoint, parameters.Openapi
	lo.ForEach(endpoints, func(item Controller, _ int) {
		item.RegisterRoutes(openapi)
	})
}
