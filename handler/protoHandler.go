package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

type ProtoHandler struct {
	UsernamePasswordHandler *UsernamePasswordProtoHandler
}

func NewProtoHandler(
	usernamePasswordHandler *UsernamePasswordProtoHandler,
) *ProtoHandler {
	return &ProtoHandler{
		UsernamePasswordHandler: usernamePasswordHandler,
	}
}

func (h *ProtoHandler) HandleRequest(data []byte) []byte {
	wrapper := &messages.RequestWrapper{}
	err := proto.Unmarshal(data, wrapper)

	response := &messages.ResponseWrapper{Ok: true}

	if err != nil {
		log.WithField("error", err).Error(ErrorUnhandledRequestMessage)

		decorateErrorResponse(response, ErrorUnhandledRequestCode, ErrorUnhandledRequestMessage)
		bytes, _ := proto.Marshal(response)

		return bytes
	}

	switch rType := wrapper.GetRequestType(); rType {
	case "UsernamePassword":
		log.WithField("type", rType).Info("Received UsernamePassword authentication request")

		h.UsernamePasswordHandler.HandleRequest(wrapper, response)
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

func decorateErrorResponse(
	response *messages.ResponseWrapper,
	errorCode uint32,
	errorMessage string,
) {
	response.Ok = false
	response.ErrorCode = errorCode
	response.ErrorMessage = errorMessage
}
