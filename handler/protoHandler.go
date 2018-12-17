package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

// HandleProtoRequest handle
func HandleProtoRequest(data []byte) []byte {
	var wrapper messages.RequestWrapper
	err := proto.Unmarshal(data, &wrapper)

	response := &messages.ResponseWrapper{Ok: true}

	if err != nil {
		log.WithField("error", err).Error(ErrorUnhandledRequestMessage)

		response.Ok = false
		response.ErrorCode = ErrorUnhandledRequestCode
		response.ErrorMessage = ErrorUnhandledRequestMessage

		responseBytes, _ := proto.Marshal(response)

		return responseBytes
	}

	switch rType := wrapper.GetRequestType(); rType {
	case "UsernamePassword":
		log.WithField("type", rType).Info("Received UsernamePassword authentication request")
		// TODO handle messages.UsernamePasswordRequestPayload
		break
	default:
		log.WithField("type", rType).Warn(ErrorUnknownRequestTypeMessage)
		response.Ok = false
		response.ErrorCode = ErrorUnknownRequestTypeCode
		response.ErrorMessage = ErrorUnknownRequestTypeMessage
		break
	}

	responseBytes, _ := proto.Marshal(response)

	return responseBytes
}
