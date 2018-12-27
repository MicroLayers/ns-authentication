package service_test

import (
	"ns-auth/service"
	"ns-auth/storage"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestMemoryStorage_UsernamePassword_AddUser(t *testing.T) {
	t.Parallel()

	hasher := service.NewStdHasher(service.HasherSalt(uuid.New().String()))
	service := service.NewMemoryStorage(hasher)

	username := uuid.New().String()
	domain := uuid.New().String()
	password := uuid.New().String()

	user, err := service.UsernamePassword.AddUser(username, domain, password)

	assert.NoError(t, err)
	assert.NotEqual(t, "", user.ID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, domain, user.Domain)
}

func TestMemoryStorage_UsernamePassword_FindUser(t *testing.T) {
	t.Parallel()

	hasher := service.NewStdHasher(service.HasherSalt(uuid.New().String()))
	service := service.NewMemoryStorage(hasher)

	username := uuid.New().String()
	domain := uuid.New().String()
	password := uuid.New().String()

	createdUser, err := service.UsernamePassword.AddUser(username, domain, password)
	assert.NoError(t, err)

	user, err := service.UsernamePassword.FindUser(username, domain, password)
	assert.NoError(t, err)
	assert.Equal(t, createdUser.ID, user.ID)
	assert.Equal(t, username, user.Username)
	assert.Equal(t, domain, user.Domain)
}

func TestMemoryStorage_Token_FindOrCreateTokenFromUser_and_FindUserFromToken(t *testing.T) {
	t.Parallel()

	hasher := service.NewStdHasher(service.HasherSalt(uuid.New().String()))
	service := service.NewMemoryStorage(hasher)

	username := uuid.New().String()
	domain := uuid.New().String()
	password := uuid.New().String()

	user, err := service.UsernamePassword.AddUser(username, domain, password)
	assert.NoError(t, err)

	authToken, err := service.Token.FindOrCreateTokenFromUser(user, storage.AuthTypeUsernamePassword)
	assert.NoError(t, err)
	assert.NotNil(t, authToken)
	assert.NotEqual(t, "", authToken.Token)
	assert.Equal(t, "", authToken.RefreshToken)
	assert.NotEqual(t, 0, authToken.ExpiryDate)

	foundUser, err := service.Token.FindUserFromToken(authToken.Token)
	assert.NoError(t, err)
	assert.Equal(t, user.ID, foundUser.ID)
}
