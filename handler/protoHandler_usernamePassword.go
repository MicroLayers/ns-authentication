package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

func handleUsernamePasswordRequest(
	wrapper *messages.RequestWrapper,
	response *messages.ResponseWrapper,
) {
	var payload messages.UsernamePasswordRequestPayload
	err := proto.Unmarshal(wrapper.GetPayload(), &payload)

	if err != nil {
		log.WithField("error", err).Error(ErrorPayloadUnmarshalMessage)
		decorateErrorResponse(response, ErrorPayloadUnmarshalCode, ErrorPayloadUnmarshalMessage)
	}

	// TODO handle messages.UsernamePasswordRequestPayload

	response.Ok = true
}
