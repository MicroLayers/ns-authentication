package handler_test

import (
	"ns-auth/handler"
	"ns-auth/messages"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func createUser(
	t *testing.T,
	protoHandler *handler.ProtoHandler,
	username string,
	password string,
	domain string,
	shouldSuccess bool,
) {
	payload := &messages.UsernamePasswordAddUserRequestPayload{
		Username: username,
		Password: password,
		Domain:   domain,
	}

	payloadBytes, err := proto.Marshal(payload)
	assert.NoError(t, err)

	message := &messages.RequestWrapper{
		RequestType: "UsernamePasswordAddUser",
		Payload:     payloadBytes,
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := protoHandler.HandleRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)

	if shouldSuccess {
		assert.True(t, data.Ok, "The handler should correctly create a new user")
	} else {
		assert.False(t, data.Ok, "The AddUser should fail")
	}
}

func TestHandleProto_UsernamePasswordProtoHandler_HandleAuthenticationRequest_willSuccess(t *testing.T) {
	t.Parallel()

	username := "username"
	password := "password"
	domain := "domain"

	protoHandler := handler.GetProtoHandler(getMemoryConfiguration())

	createUser(t, protoHandler, username, password, domain, true)

	payload := &messages.UsernamePasswordLoginRequestPayload{
		Username: username,
		Password: password,
		Domain:   domain,
	}

	payloadBytes, err := proto.Marshal(payload)
	assert.NoError(t, err)

	message := &messages.RequestWrapper{
		RequestType: "UsernamePasswordAuthentication",
		Payload:     payloadBytes,
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := protoHandler.HandleRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)
	assert.True(t, data.Ok, "The handler should process a right payload")
}

func TestHandleProto_UsernamePasswordProtoHandler_HandleAuthenticationRequest_wrongRequestPayload(t *testing.T) {
	t.Parallel()

	protoHandler := handler.GetProtoHandler(getMemoryConfiguration())

	payloadBytes := []byte("some random bytes here")

	message := &messages.RequestWrapper{
		RequestType: "UsernamePasswordAuthentication",
		Payload:     payloadBytes,
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := protoHandler.HandleRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)
	assert.False(t, data.Ok, "The handler should reject a wrong payload")
}

func TestHandleProto_UsernamePasswordProtoHandler_HandleAddUserRequest_shouldCorrectlyThrowAnErrorTryingToCreateAUserTwoTimes(t *testing.T) {
	t.Parallel()

	username := "username"
	password := "password"
	domain := "domain"

	protoHandler := handler.GetProtoHandler(getMemoryConfiguration())

	createUser(t, protoHandler, username, password, domain, true)
	createUser(t, protoHandler, username, password, domain, false)
}
