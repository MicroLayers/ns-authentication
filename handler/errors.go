package handler

// ErrorUnhandledRequestCode unable to parse the request
const ErrorUnhandledRequestCode = uint32(3000)

// ErrorUnhandledRequestMessage unable to parse the request
const ErrorUnhandledRequestMessage = "Unable to parse the request"

// ErrorUnknownRequestTypeCode unknown request type received
const ErrorUnknownRequestTypeCode = uint32(3001)

// ErrorUnknownRequestTypeMessage unknown request type received
const ErrorUnknownRequestTypeMessage = "Unknown request type received"

// ErrorPayloadUnmarshalCode invalid request payload (unable to parse)
const ErrorPayloadUnmarshalCode = uint32(3002)

// ErrorPayloadUnmarshalMessage invalid request payload (unable to parse)
const ErrorPayloadUnmarshalMessage = "Unable to parse the request's payload"

// ErrorInvalidCredentialsCode the supplied credentials are invalid
const ErrorInvalidCredentialsCode = uint32(3003)

// ErrorInvalidCredentialsMessage the supplied credentials are invalid
const ErrorInvalidCredentialsMessage = "Invalid credentials received, user not found or invalid password"

// ErrorUnableToGenerateResponsePayloadCode unable to marshal the response
const ErrorUnableToGenerateResponsePayloadCode = uint32(3004)

// ErrorUnableToGenerateResponsePayloadMessage unable to marshal the response
const ErrorUnableToGenerateResponsePayloadMessage = "Unable to generate the response's payload"

// ErrorUnableToCreateNewUserCode unable to create a new user
const ErrorUnableToCreateNewUserCode = uint32(3005)

// ErrorUnableToCreateNewUserMessage unable to create a new user
const ErrorUnableToCreateNewUserMessage = "Unable to create a new user"

// ErrorUnableToFindUserFromTokenAndDomainCode unable to find a user for the given token and domain
const ErrorUnableToFindUserFromTokenAndDomainCode = uint32(3006)

// ErrorUnableToFindUserFromTokenAndDomainMessage unable to find a user for the given token and domain
const ErrorUnableToFindUserFromTokenAndDomainMessage = "Unable to find a user for the given token and domain"
