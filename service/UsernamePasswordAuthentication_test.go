package service_test

import (
	"ns-auth/configuration"
	"ns-auth/service"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func getMemoryUsernamePasswordAuth() *service.UsernamePasswordAuthentication {
	config := &configuration.Configuration{}
	config.Authentication.Store.Type = "memory"

	return service.GetUsernamePasswordAuthentication(config)
}

func TestUsernamePasswordAuthentication_AddUser(t *testing.T) {
	t.Parallel()

	username := uuid.New().String()
	password := uuid.New().String()
	domain := uuid.New().String()

	service := getMemoryUsernamePasswordAuth()

	user, err := service.AddUser(username, password, domain)
	assert.NoError(t, err)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, domain, user.Domain)
	assert.NotEmpty(t, user.ID)
}

func TestUsernamePasswordAuthentication_GetAuthToken(t *testing.T) {
	t.Parallel()

	username := uuid.New().String()
	password := uuid.New().String()
	domain := uuid.New().String()

	service := getMemoryUsernamePasswordAuth()

	_, err := service.AddUser(username, password, domain)
	assert.NoError(t, err)

	token, err := service.GetAuthToken(username, password, domain)

	assert.NoError(t, err)
	assert.NotEmpty(t, token.Token)
	assert.Empty(t, token.RefreshToken)
	assert.NotEqual(t, int64(0), token.ExpiryDate)
}
