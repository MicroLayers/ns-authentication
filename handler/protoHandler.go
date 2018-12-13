package handler

import (
	"ns-auth/messages"

	"github.com/golang/protobuf/proto"
	log "github.com/sirupsen/logrus"
)

// HandleProtoRequest handle
func HandleProtoRequest(data []byte) []byte {
	var baseRequest messages.BaseRequest

	err := proto.Unmarshal(data, &baseRequest)

	if err != nil {
		log.WithField("error", err).Error("Received unhandled request")
		// TODO manage error
		return data
	}

	switch rType := baseRequest.GetRequestType(); rType {
	case "UsernamePassword":
		log.WithField("type", rType).Info("Received UsernamePassword authentication request")
		// TODO handle messages.UsernamePasswordAuthenticationRequest
		break
	default:
		log.WithField("type", rType).Warn("Received unhandled request type")
	}

	// TODO manage return
	return data
}
