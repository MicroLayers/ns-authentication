package service_test

import (
	"errors"
	"ns-auth/service"
	"ns-auth/storage"
	"testing"

	"github.com/stretchr/testify/assert"
)

type hasherConcat struct{}

func (h *hasherConcat) HashPassword(username string, password string, domain string) string {
	return username + password + domain
}

type usernamePasswordStorage struct {
	ExpectedUsername string
	ExpectedPassword string
	ExpectedDomain   string
	ExpectedUserId   string
}

func (ups *usernamePasswordStorage) FindUser(username string, password string, domain string) (storage.User, error) {
	sameUsername := ups.ExpectedUsername == username
	samePassword := ups.ExpectedPassword == password
	sameDomain := ups.ExpectedDomain == domain
	if sameUsername && samePassword && sameDomain {
		return storage.User{
			Id:       ups.ExpectedUserId,
			Username: username,
			Domain:   domain,
		}, nil
	}

	return storage.User{}, errors.New("")
}

type tokenStorage struct {
	ExpectedUserId         string
	ExpectedToken          string
	ExpectedRefreshToken   string
	ExpectedExpirationDate uint32
}

func (ts *tokenStorage) FindOrCreate(user storage.User, authType string) (storage.AuthToken, error) {
	if user.Id == ts.ExpectedUserId && authType == storage.AUTH_TYPE_USERNAME_PASSWORD {
		return storage.AuthToken{
			Token:          ts.ExpectedToken,
			RefreshToken:   ts.ExpectedRefreshToken,
			ExpirationDate: ts.ExpectedExpirationDate,
		}, nil
	}

	return storage.AuthToken{}, errors.New("")
}

func TestGetAuthToken(t *testing.T) {
	username := "username"
	password := "password"
	domain := "domain"

	expectedUserId := "whatever_user"
	expectedToken := "0123456789abcdef"
	expectedRefreshToken := "fedcba9876543210"
	expectedExpirationDate := uint32(1234567)

	hasher := hasherConcat{}
	storage := storage.Storage{
		UsernamePassword: &usernamePasswordStorage{
			ExpectedUserId:   expectedUserId,
			ExpectedUsername: username,
			ExpectedPassword: hasher.HashPassword(username, password, domain),
			ExpectedDomain:   domain,
		},
		Token: &tokenStorage{
			ExpectedUserId:         expectedUserId,
			ExpectedToken:          expectedToken,
			ExpectedRefreshToken:   expectedRefreshToken,
			ExpectedExpirationDate: expectedExpirationDate,
		},
	}

	service := service.UsernamePasswordAuthentication{
		Hasher:  &hasher,
		Storage: storage,
	}

	token, err := service.GetAuthToken(username, password, domain)

	assert.NoError(t, err)
	assert.Equal(t, expectedToken, token.Token)
	assert.Equal(t, expectedRefreshToken, token.RefreshToken)
	assert.Equal(t, expectedExpirationDate, token.ExpirationDate)
}
