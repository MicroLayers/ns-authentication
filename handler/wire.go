//+build wireinject

package handler

import (
	"ns-auth/configuration"
	"ns-auth/service"

	"github.com/google/wire"
)

// GetUsernamePasswordAuthentication service provider
func GetProtoHandler(
	config *configuration.Configuration,
) *ProtoHandler {
	wire.Build(
		NewProtoHandler,
		NewUsernamePasswordProtoHandler,
		service.GetUsernamePasswordAuthentication,
	)

	return &ProtoHandler{}
}
