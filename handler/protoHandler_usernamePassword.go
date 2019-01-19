package handler

import (
	"fmt"
	"ns-auth/messages"
	"ns-auth/service"

	"github.com/golang/protobuf/proto"
)

// UsernamePasswordProtoHandler the Protobuf handler used to manage username/password related requests
type UsernamePasswordProtoHandler struct {
	usernamePasswordService *service.UsernamePasswordAuthentication
}

// NewUsernamePasswordProtoHandler UsernamePasswordProtoHandler's instantiator used by wire
func NewUsernamePasswordProtoHandler(
	usernamePasswordService *service.UsernamePasswordAuthentication,
) *UsernamePasswordProtoHandler {
	return &UsernamePasswordProtoHandler{
		usernamePasswordService: usernamePasswordService,
	}
}

// HandleAuthenticationRequest handle the authentication request
func (h *UsernamePasswordProtoHandler) HandleAuthenticationRequest(
	wrapper *messages.RequestWrapper,
	response *messages.ResponseWrapper,
) {
	var payload messages.UsernamePasswordLoginRequestPayload
	err := unmarshalPayloadOrError(response, wrapper, &payload)
	if err != nil {
		// Error handled in previous function
		return
	}

	username := payload.GetUsername()
	password := payload.GetPassword()
	domain := payload.GetDomain()

	token, err := h.usernamePasswordService.GetAuthToken(
		username,
		password,
		domain,
	)
	if err != nil {
		logAndDecorateNegativeResponse(
			response,
			ErrorInvalidCredentialsCode,
			ErrorInvalidCredentialsMessage,
			err,
		)

		return
	}

	responsePayload := &messages.AuthenticationResponse{
		AuthToken:    token.Token,
		RefreshToken: token.RefreshToken,
		ExpiryDate:   token.ExpiryDate,
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
		fmt.Sprintf("Succesfully authenticated user %s - %s", domain, password),
	)
}

// HandleAddUserRequest handle the add user request
func (h *UsernamePasswordProtoHandler) HandleAddUserRequest(
	wrapper *messages.RequestWrapper,
	response *messages.ResponseWrapper,
) {
	var payload messages.UsernamePasswordAddUserRequestPayload
	err := unmarshalPayloadOrError(response, wrapper, &payload)
	if err != nil {
		// Error handled in previous function
		return
	}

	user, err := h.usernamePasswordService.AddUser(
		payload.GetUsername(),
		payload.GetPassword(),
		payload.GetDomain(),
	)

	if err != nil {
		logAndDecorateNegativeResponse(
			response,
			ErrorUnableToCreateNewUserCode,
			ErrorUnableToCreateNewUserMessage,
			err,
		)

		return
	}

	responsePayload := &messages.UserCreationResponsePayload{
		ID: user.ID,
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
		fmt.Sprintf("Created user %s - %s", user.Domain, user.Username),
	)
}
