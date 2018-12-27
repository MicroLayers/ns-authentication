package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

// ProtoHandler Protobuf handler
type ProtoHandler struct {
	UsernamePasswordHandler *UsernamePasswordProtoHandler
}

// NewProtoHandler ProtoHandler's instantiator used by wire
func NewProtoHandler(
	usernamePasswordHandler *UsernamePasswordProtoHandler,
) *ProtoHandler {
	return &ProtoHandler{
		UsernamePasswordHandler: usernamePasswordHandler,
	}
}

// HandleRequest handle the proto requests
func (h *ProtoHandler) HandleRequest(data []byte) []byte {
	wrapper := &messages.RequestWrapper{}
	err := proto.Unmarshal(data, wrapper)

	response := &messages.ResponseWrapper{Ok: true}

	if err != nil {
		logAndDecorateNegativeResponse(response, ErrorUnhandledRequestCode, ErrorUnhandledRequestMessage, err)
		bytes, _ := proto.Marshal(response)

		return bytes
	}

	switch rType := wrapper.GetRequestType(); rType {
	case "UsernamePasswordAuthentication":
		log.WithField("type", rType).Info("Received UsernamePassword authentication request")

		h.UsernamePasswordHandler.HandleAuthenticationRequest(wrapper, response)
		break
	case "UsernamePasswordAddUser":
		log.WithField("type", rType).Info("Received UsernamePassword add user request")
		h.UsernamePasswordHandler.HandleAddUserRequest(wrapper, response)
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

func logAndDecorateNegativeResponse(
	response *messages.ResponseWrapper,
	errorCode uint32,
	errorMessage string,
	err error,
) {
	log.WithField("error", err).Error(ErrorPayloadUnmarshalMessage)

	response.Ok = false
	response.ErrorCode = errorCode
	response.ErrorMessage = errorMessage
}

func logAndDecoratePositiveResponse(
	response *messages.ResponseWrapper,
	payload []byte,
	messageToLog string,
) {
	log.Info(messageToLog)

	response.Ok = true
	response.Payload = payload
}
