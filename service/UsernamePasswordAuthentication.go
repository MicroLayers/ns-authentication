package service

import (
	"ns-auth/storage"
)

type UsernamePasswordAuthentication struct {
	Storage *storage.Storage
	Hasher  storage.Hasher
}

func NewUsernamePasswordAuthentication(
	storage *storage.Storage,
	hasher storage.Hasher,
) *UsernamePasswordAuthentication {
	checker := UsernamePasswordAuthentication{Storage: storage, Hasher: hasher}

	return &checker
}

func (c *UsernamePasswordAuthentication) AddUser(
	username string,
	password string,
	domain string,
) (*storage.User, error) {
	// User already exists
	if user, err := c.Storage.UsernamePassword.FindUser(username, domain, password); err == nil {
		return user, nil
	}

	return c.Storage.UsernamePassword.AddUser(username, domain, password)
}

func (c *UsernamePasswordAuthentication) GetAuthToken(
	username string,
	password string,
	domain string,
) (authToken *storage.AuthToken, err error) {
	user, err := c.Storage.UsernamePassword.FindUser(username, domain, password)

	if err != nil {
		return authToken, err
	}

	return c.Storage.Token.FindOrCreateTokenFromUser(user, storage.AUTH_TYPE_USERNAME_PASSWORD)
}
