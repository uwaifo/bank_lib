package errs

import "net/http"

type AppError struct {
	Code    int `json:",omitempty"`
	Message string
}

func (e AppError) AsMessage() *AppError {
	return &AppError{
		Message: e.Message,
	}
}

// NOTICE * Declared returned type
func NewNotFoundError(message string) *AppError {

	//NOTICE Returns & of pointer
	return &AppError{
		Code:    http.StatusNotFound,
		Message: message,
	}

}

func NewUnexpectedError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusInternalServerError,
		Message: msg,
	}
}

func NewValidationError(msg string) *AppError {
	return &AppError{
		Code:    http.StatusUnprocessableEntity, // 422
		Message: msg,
	}
}

func NewAuthenticationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusUnauthorized,
	}
}

func NewAuthorizationError(message string) *AppError {
	return &AppError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
