package response

import "github.com/jfavo/integra-partners-assessment-backend/internal/errors"

type Response struct {
	Data         interface{}      `json:"data,omitempty"`
	ErrorCode    errors.ErrorCode `json:"error_code,omitempty"`
	ErrorMessage string           `json:"error_message,omitempty"`
}

// Success returns a successful response object to the user containing
// the data requested.
func Success(data interface{}) Response {
	return Response{
		Data: data,
	}
}

// Failure returns a non-successful response obejct to the user containing
// the error code and message to the client.
func Failure(errCode errors.ErrorCode, errMessage string) Response {
	return Response{
		ErrorCode:    errCode,
		ErrorMessage: errMessage,
	}
}
