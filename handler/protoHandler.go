package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

const RequestTypeTokenDiscover string = "TokenDiscover"
const RequestTypeUsernamePasswordAuthentication string = "UsernamePasswordAuthentication"
const RequestTypeUsernamePasswordAddUser string = "UsernamePasswordAddUser"

// ProtoHandler Protobuf handler
type ProtoHandler struct {
	UsernamePasswordHandler *UsernamePasswordProtoHandler
	TokenHandler            *TokenProtoHandler
}

// NewProtoHandler ProtoHandler's instantiator used by wire
func NewProtoHandler(
	usernamePasswordHandler *UsernamePasswordProtoHandler,
	tokenHandler *TokenProtoHandler,
) *ProtoHandler {
	return &ProtoHandler{
		UsernamePasswordHandler: usernamePasswordHandler,
		TokenHandler:            tokenHandler,
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
	case RequestTypeUsernamePasswordAuthentication:
		log.WithField("type", rType).Info("Received UsernamePassword authentication request")

		h.UsernamePasswordHandler.HandleAuthenticationRequest(wrapper, response)
		break
	case RequestTypeUsernamePasswordAddUser:
		log.WithField("type", rType).Info("Received UsernamePassword add user request")
		h.UsernamePasswordHandler.HandleAddUserRequest(wrapper, response)
	case RequestTypeTokenDiscover:
		log.WithField("type", rType).Info("Received TokenDiscover request")
		h.TokenHandler.HandleTokenDiscoverRequest(wrapper, response)
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

func unmarshalPayloadOrError(
	response *messages.ResponseWrapper,
	wrapper *messages.RequestWrapper,
	dest proto.Message,
) error {
	err := proto.Unmarshal(wrapper.GetPayload(), dest)

	if err != nil {
		logAndDecorateNegativeResponse(
			response,
			ErrorPayloadUnmarshalCode,
			ErrorPayloadUnmarshalMessage,
			err,
		)

		return err
	}

	return nil
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
