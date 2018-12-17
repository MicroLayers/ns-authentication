package handler_test

import (
	"ns-auth/handler"
	"ns-auth/messages"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestHandleProto_UsernamePasswordRequest(t *testing.T) {
	payload := &messages.UsernamePasswordRequestPayload{
		Username: "Username",
		Password: "Password",
	}

	payloadBytes, err := proto.Marshal(payload)
	assert.NoError(t, err)

	message := &messages.RequestWrapper{
		RequestType: "UsernamePassword",
		Payload:     payloadBytes,
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := handler.HandleProtoRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)
	assert.True(t, data.Ok)
}

func TestHandleProto_UnknownRequestType(t *testing.T) {
	message := &messages.RequestWrapper{
		RequestType: "Unknown type",
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := handler.HandleProtoRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)
	assert.False(t, data.Ok)
	assert.Equal(t, handler.ErrorUnknownRequestTypeCode, data.ErrorCode)
	assert.Equal(t, handler.ErrorUnknownRequestTypeMessage, data.ErrorMessage)
}
