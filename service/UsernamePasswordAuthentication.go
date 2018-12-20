package service

import (
	"ns-auth/storage"
)

type UsernamePasswordAuthentication struct {
	Storage storage.Storage
	Hasher  storage.Hasher
}

func NewUsernamePasswordAuthentication(storage storage.Storage, hasher storage.Hasher) *UsernamePasswordAuthentication {
	checker := UsernamePasswordAuthentication{Storage: storage, Hasher: hasher}

	return &checker
}

func (c *UsernamePasswordAuthentication) GetAuthToken(username string, password string, domain string) (authToken storage.AuthToken, err error) {
	hashedPassword := c.Hasher.HashPassword(username, password, domain)
	user, err := c.Storage.UsernamePassword.FindUser(username, hashedPassword, domain)

	if err != nil {
		return authToken, err
	}
	return c.Storage.Token.FindOrCreate(user, storage.AUTH_TYPE_USERNAME_PASSWORD)
}
