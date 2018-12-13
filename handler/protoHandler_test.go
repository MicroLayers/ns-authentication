package handler_test

import (
	"ns-auth/handler"
	"ns-auth/messages"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestHandleProto(t *testing.T) {
	message := &messages.UsernamePasswordAuthenticationRequest{
		RequestType: "UsernamePassword",
		Username:    "Username",
		Password:    "Password",
	}

	buffer, err := proto.Marshal(message)

	assert.NoError(t, err)

	handler.HandleProtoRequest(buffer)
}
