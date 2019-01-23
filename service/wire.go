//+build wireinject

package service

import (
	"ns-auth/configuration"
	"ns-auth/storage"
	"sync"

	"github.com/google/wire"
)

var storageInstanceMutex sync.RWMutex
var storageInstance *storage.Storage

func GetStorage(
	config *configuration.Configuration,
	hasher storage.Hasher,
) *storage.Storage {
	storageInstanceMutex.RLock()

	if storageInstance != nil {
		storageInstanceMutex.RUnlock()

		return storageInstance
	}
	storageInstanceMutex.RUnlock()

	var instance *storage.Storage

	switch config.Authentication.Store.Type {
	case "memory":
		instance = NewMemoryStorage(hasher)
	default:
		instance = NewMemoryStorage(hasher)
	}

	storageInstanceMutex.Lock()
	if instance != nil {
		storageInstance = instance
	}
	storageInstanceMutex.Unlock()

	return instance
}

func GetHasher(config *configuration.Configuration) storage.Hasher {
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
		GetStorage,
		GetHasher,
	)

	return &UsernamePasswordAuthentication{}
}
