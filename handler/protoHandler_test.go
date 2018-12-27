package handler_test

import (
	"ns-auth/handler"
	"ns-auth/messages"
	"testing"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestHandleProto_UnknownRequestType(t *testing.T) {
	protoHandler := handler.NewProtoHandler(&handler.UsernamePasswordProtoHandler{})

	message := &messages.RequestWrapper{
		RequestType: "Unknown type",
	}

	buffer, err := proto.Marshal(message)
	assert.NoError(t, err)

	response := protoHandler.HandleRequest(buffer)

	var data messages.ResponseWrapper
	err = proto.Unmarshal(response, &data)
	assert.NoError(t, err)
	assert.False(t, data.Ok)
	assert.Equal(t, handler.ErrorUnknownRequestTypeCode, data.ErrorCode)
	assert.Equal(t, handler.ErrorUnknownRequestTypeMessage, data.ErrorMessage)
}
