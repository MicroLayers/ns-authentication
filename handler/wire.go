//+build wireinject

package handler

import (
	"ns-auth/configuration"

	"github.com/google/wire"
)

// GetUsernamePasswordAuthentication service provider
func GetProtoHandler(
	config *configuration.Configuration,
) *ProtoHandler {
	wire.Build(
		NewProtoHandler,
		NewUsernamePasswordProtoHandler,
	)

	return &ProtoHandler{}
}
