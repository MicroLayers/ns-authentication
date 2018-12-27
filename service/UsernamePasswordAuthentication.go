package service

import (
	"fmt"
	"ns-auth/storage"
)

// UsernamePasswordAuthentication username/password service
type UsernamePasswordAuthentication struct {
	Storage *storage.Storage
	Hasher  storage.Hasher
}

// NewUsernamePasswordAuthentication UsernamePasswordAuthentication instantiator used by wire
func NewUsernamePasswordAuthentication(
	storage *storage.Storage,
	hasher storage.Hasher,
) *UsernamePasswordAuthentication {
	checker := UsernamePasswordAuthentication{Storage: storage, Hasher: hasher}

	return &checker
}

// AddUser add a new user inside the database
func (c *UsernamePasswordAuthentication) AddUser(
	username string,
	password string,
	domain string,
) (*storage.User, error) {
	// User already exists
	if user, err := c.Storage.UsernamePassword.FindUser(username, domain, password); err == nil {
		return nil, fmt.Errorf("User already exists (%s)", user.ID)
	}

	return c.Storage.UsernamePassword.AddUser(username, domain, password)
}

// GetAuthToken get an authentication token from the given credentials
func (c *UsernamePasswordAuthentication) GetAuthToken(
	username string,
	password string,
	domain string,
) (authToken *storage.AuthToken, err error) {
	user, err := c.Storage.UsernamePassword.FindUser(username, domain, password)

	if err != nil {
		return authToken, err
	}

	return c.Storage.Token.FindOrCreateTokenFromUser(user, storage.AuthTypeUsernamePassword)
}
