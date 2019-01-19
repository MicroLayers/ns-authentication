package handler

import (
	"errors"
	"ns-auth/messages"
	"ns-auth/storage"

	"github.com/golang/protobuf/proto"
)

type TokenProtoHandler struct {
	tokenStorage storage.TokenStorage
}

func NewTokenProtoHandler(tokenStorage storage.TokenStorage) *TokenProtoHandler {
	return &TokenProtoHandler{
		tokenStorage: tokenStorage,
	}
}

// HandleTokenDiscoverRequest handle the TokenDiscover request
func (h *TokenProtoHandler) HandleTokenDiscoverRequest(
	wrapper *messages.RequestWrapper,
	response *messages.ResponseWrapper,
) {
	var payload messages.TokenDiscoverRequestPayload
	err := unmarshalPayloadOrError(response, wrapper, &payload)

	if err != nil {
		// Error handled in previous function
		return
	}

	user, err := h.tokenStorage.FindUserFromToken(payload.GetToken())

	if err == nil {
		payloadDomain := payload.GetDomain()
		if user == nil || user.Domain != payloadDomain {
			err = errors.New("Token and given domains mismatch")
		}
	}

	if err != nil {
		logAndDecorateNegativeResponse(
			response,
			ErrorUnableToFindUserFromTokenAndDomainCode,
			ErrorUnableToFindUserFromTokenAndDomainMessage,
			err,
		)

		return
	}

	responsePayload := &messages.TokenDiscoverResponsePayload{
		UserID:   user.ID,
		Domain:   user.Domain,
		Username: user.Username,
	}
	responsePayloadBytes, err := proto.Marshal(responsePayload)
	if err != nil {
		logAndDecorateNegativeResponse(
			response,
			ErrorUnableToGenerateResponsePayloadCode,
			ErrorUnableToGenerateResponsePayloadMessage,
			err,
		)

		return
	}

	logAndDecoratePositiveResponse(
		response,
		responsePayloadBytes,
		"log",
	)
}
