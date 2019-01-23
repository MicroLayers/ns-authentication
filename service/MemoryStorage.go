package service

import (
	"errors"
	"ns-auth/storage"
	"sync"
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
	mutex  sync.RWMutex
}

type tokenMemoryStorage struct {
	hasher storage.Hasher
	store  *userDefinitionList
	mutex  sync.RWMutex
}

// NewMemoryStorage MemoryStorage's instantiator, used by wire
func NewMemoryStorage(hasher storage.Hasher) *storage.Storage {
	commonStore := &userDefinitionList{}

	usernamePasswordStore := &usernamePasswordMemoryStorage{
		hasher: hasher,
		store:  commonStore,
		mutex:  sync.RWMutex{},
	}

	tokenStore := &tokenMemoryStorage{
		hasher: hasher,
		store:  commonStore,
		mutex:  sync.RWMutex{},
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
	s.mutex.RLock()
	for _, definition := range *s.store {
		if s.hasher.CheckPassword(username, domain, password, definition.HashedPassword) {
			s.mutex.RUnlock()

			return definition.User, nil
		}
	}
	s.mutex.RUnlock()

	expiryDate := time.Now().Add(time.Hour).Unix() // One hour expiration

	newDef := userDefinition{
		HashedPassword: s.hasher.HashPassword(username, domain, password),
		User: &storage.User{
			ID:       uuid.New().String(),
			Username: username,
			Domain:   domain,
		},
		AuthToken: &storage.AuthToken{
			Token:        uuid.New().String(),
			RefreshToken: uuid.New().String(),
			ExpiryDate:   expiryDate,
		},
	}

	s.mutex.Lock()
	*s.store = append(*s.store, newDef)
	s.mutex.Unlock()

	return newDef.User, nil
}

func (s *usernamePasswordMemoryStorage) FindUser(
	username string,
	domain string,
	password string,
) (*storage.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	nowEpoch := time.Now().Unix()

	for _, definition := range *s.store {
		if s.hasher.CheckPassword(username, domain, password, definition.HashedPassword) &&
			nowEpoch <= definition.AuthToken.ExpiryDate {

			return definition.User, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *tokenMemoryStorage) FindOrCreateTokenFromUser(
	user *storage.User,
	authType string,
) (*storage.AuthToken, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	// ignore authType for this storage type
	for _, definition := range *s.store {
		if definition.User.ID == user.ID {
			token := definition.AuthToken
			token.RefreshToken = "" // not a new token, do not send the refresh one

			return token, nil
		}
	}

	return nil, errors.New("user not found")
}

func (s *tokenMemoryStorage) FindUserFromToken(token string) (*storage.User, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	for _, definition := range *s.store {
		if definition.AuthToken.Token == token {

			return definition.User, nil
		}
	}

	return nil, errors.New("user not found")
}
