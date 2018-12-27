package service

import (
	"errors"
	"ns-auth/storage"
	"time"

	"github.com/google/uuid"
)

// Memory storage is used by tests to avoid using any kind of database

type userDefinition struct {
	HashedPassword string
	User           *storage.User
	AuthToken      *storage.AuthToken
}

type userDefinitionList []userDefinition

type usernamePasswordMemoryStorage struct {
	hasher storage.Hasher
	store  *userDefinitionList
}

type tokenMemoryStorage struct {
	hasher storage.Hasher
	store  *userDefinitionList
}

func NewMemoryStorage(hasher storage.Hasher) *storage.Storage {
	commonStore := &userDefinitionList{}

	usernamePasswordStore := &usernamePasswordMemoryStorage{
		hasher: hasher,
		store:  commonStore,
	}

	tokenStore := &tokenMemoryStorage{
		hasher: hasher,
		store:  commonStore,
	}

	return &storage.Storage{
		UsernamePassword: usernamePasswordStore,
		Token:            tokenStore,
	}
}

func (s *usernamePasswordMemoryStorage) AddUser(
	username string,
	domain string,
	password string,
) (*storage.User, error) {
	for _, definition := range *s.store {
		if s.hasher.CheckPassword(username, domain, password, definition.HashedPassword) {
			return definition.User, nil
		}
	}

	expiration := time.Now().Add(time.Hour).Unix()

	newDef := userDefinition{
		HashedPassword: s.hasher.HashPassword(username, domain, password),
		User: &storage.User{
			Id:       uuid.New().String(),
			Username: username,
			Domain:   domain,
		},
		AuthToken: &storage.AuthToken{
			Token:          uuid.New().String(),
			RefreshToken:   uuid.New().String(),
			ExpirationDate: uint32(expiration),
		},
	}
	*s.store = append(*s.store, newDef)

	return newDef.User, nil
}

func (s *usernamePasswordMemoryStorage) FindUser(
	username string,
	domain string,
	password string,
) (*storage.User, error) {
	for _, definition := range *s.store {
		if s.hasher.CheckPassword(username, domain, password, definition.HashedPassword) {
			return definition.User, nil
		}
	}

	return nil, errors.New("User not found!")
}

func (s *tokenMemoryStorage) FindOrCreateTokenFromUser(
	user *storage.User,
	authType string,
) (*storage.AuthToken, error) {
	// ignore authType for this storage type
	for _, definition := range *s.store {
		if definition.User.Id == user.Id {
			return definition.AuthToken, nil
		}
	}

	return nil, errors.New("User not found!")
}

func (s *tokenMemoryStorage) FindUserFromToken(token string) (*storage.User, error) {
	for _, definition := range *s.store {
		if definition.AuthToken.Token == token {
			return definition.User, nil
		}
	}

	return nil, errors.New("User not found!")
}
