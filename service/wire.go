//+build wireinject

package service

import (
	"ns-auth/configuration"
	"ns-auth/storage"

	"github.com/google/wire"
)

func _injectorStorage(
	config *configuration.Configuration,
	hasher storage.Hasher,
) *storage.Storage {
	switch config.Authentication.Store.Type {
	case "memory":
	default:
		return NewMemoryStorage(hasher)
	}

	// wire workaround
	return NewMemoryStorage(hasher)
}

func _injectorHasher(config *configuration.Configuration) storage.Hasher {
	salt := HasherSalt(config.Authentication.Hasher.Salt)

	switch config.Authentication.Hasher.Type {
	case "standard":
	default:
		return NewStdHasher(salt)
	}

	// wire workaround
	return NewStdHasher(salt)
}

// GetUsernamePasswordAuthentication service provider
func GetUsernamePasswordAuthentication(
	config *configuration.Configuration,
) *UsernamePasswordAuthentication {
	wire.Build(
		NewUsernamePasswordAuthentication,
		_injectorHasher,
		_injectorStorage,
	)

	return &UsernamePasswordAuthentication{}
}
