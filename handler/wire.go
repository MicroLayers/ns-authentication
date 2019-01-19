//+build wireinject

package handler

import (
	"ns-auth/configuration"
	"ns-auth/service"
	"ns-auth/storage"

	"github.com/google/wire"
)

func _tokenStorageInjector(
	storage *storage.Storage,
) storage.TokenStorage {
	return storage.Token
}

// GetUsernamePasswordAuthentication service provider
func GetProtoHandler(
	config *configuration.Configuration,
) *ProtoHandler {
	wire.Build(
		NewProtoHandler,
		NewTokenProtoHandler,
		NewUsernamePasswordProtoHandler,
		service.GetUsernamePasswordAuthentication,
		service.GetStorage,
		service.GetHasher,
		_tokenStorageInjector,
	)

	return &ProtoHandler{}
}
