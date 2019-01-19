package handler_test

import (
	"ns-auth/handler"
	"ns-auth/messages"
	"ns-auth/service"
	"ns-auth/storage"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func createUserWithPasswordViaService(
	username string,
	domain string,
	password string,
) (*storage.User, error) {
	memoryConfiguration := getMemoryConfiguration()
	hasher := service.GetHasher(memoryConfiguration)
	memoryStorage := service.GetStorage(memoryConfiguration, hasher)

	return memoryStorage.UsernamePassword.AddUser(username, domain, password)
}

func getTokenForUserWithPasswordViaService(user *storage.User) (*storage.AuthToken, error) {
	memoryConfiguration := getMemoryConfiguration()
	hasher := service.GetHasher(memoryConfiguration)
	memoryStorage := service.GetStorage(memoryConfiguration, hasher)

	return memoryStorage.Token.FindOrCreateTokenFromUser(user, storage.AuthTypeUsernamePassword)
}

func processPayloadBytes(t *testing.T, payload []byte) *messages.ResponseWrapper {
	protoHandler := handler.GetProtoHandler(getMemoryConfiguration())

	requestWrapper := &messages.RequestWrapper{
		RequestType: handler.RequestTypeTokenDiscover,
		Payload:     payload,
	}

	buffer, err := proto.Marshal(requestWrapper)
	assert.NoError(t, err)
	response := protoHandler.HandleRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)

	return &data
}

func processPayload(t *testing.T, payload proto.Message) *messages.ResponseWrapper {
	payloadBytes, err := proto.Marshal(payload)
	assert.NoError(t, err)

	return processPayloadBytes(t, payloadBytes)
}

func TestHandleProto_TokenHandler_HandleTokenDiscoverRequest_failWrongRequestPayloadBytes(
	t *testing.T,
) {
	t.Parallel()

	payload := []byte("this is not a proper payload")
	data := processPayloadBytes(t, payload)

	assert.False(t, data.Ok)
	assert.Equal(t, handler.ErrorPayloadUnmarshalCode, data.ErrorCode)
	assert.Equal(t, handler.ErrorPayloadUnmarshalMessage, data.ErrorMessage)
}

func TestHandleProto_TokenHandler_HandleTokenDiscoverRequest_failNoUserFound(
	t *testing.T,
) {
	t.Parallel()

	payload := &messages.TokenDiscoverRequestPayload{
		Token:  "myToken",
		Domain: "MyDomain",
	}

	data := processPayload(t, payload)

	assert.False(t, data.Ok)
	assert.Equal(t, handler.ErrorUnableToFindUserFromTokenAndDomainCode, data.ErrorCode)
	assert.Equal(t, handler.ErrorUnableToFindUserFromTokenAndDomainMessage, data.ErrorMessage)
}

func TestHandleProto_TokenHandler_HandleTokenDiscoverRequest_failWrongDomain(
	t *testing.T,
) {
	t.Parallel()

	username := uuid.New().String()
	password := uuid.New().String()
	domain := uuid.New().String()
	wrongDomain := uuid.New().String()

	user, err := createUserWithPasswordViaService(username, domain, password)
	assert.NoError(t, err)
	token, err := getTokenForUserWithPasswordViaService(user)
	assert.NoError(t, err)

	payload := &messages.TokenDiscoverRequestPayload{
		Token:  token.Token,
		Domain: wrongDomain,
	}

	data := processPayload(t, payload)

	assert.False(t, data.Ok)
	assert.Equal(t, handler.ErrorUnableToFindUserFromTokenAndDomainCode, data.ErrorCode)
	assert.Equal(t, handler.ErrorUnableToFindUserFromTokenAndDomainMessage, data.ErrorMessage)
}

func TestHandleProto_TokenHandler_HandleTokenDiscoverRequest_success(t *testing.T) {
	t.Parallel()

	username := uuid.New().String()
	password := uuid.New().String()
	domain := uuid.New().String()

	user, err := createUserWithPasswordViaService(username, domain, password)
	assert.NoError(t, err)
	token, err := getTokenForUserWithPasswordViaService(user)
	assert.NoError(t, err)

	payload := &messages.TokenDiscoverRequestPayload{
		Token:  token.Token,
		Domain: domain,
	}

	data := processPayload(t, payload)
	assert.True(t, data.Ok)
	assert.Empty(t, data.ErrorCode)
	assert.Empty(t, data.ErrorMessage)

	var responsePayload messages.TokenDiscoverResponsePayload
	err = proto.Unmarshal(data.Payload, &responsePayload)
	assert.NoError(t, err)

	assert.Equal(t, domain, responsePayload.GetDomain())
	assert.Equal(t, user.ID, responsePayload.GetUserID())
	assert.Equal(t, user.Username, responsePayload.GetUsername())
}
