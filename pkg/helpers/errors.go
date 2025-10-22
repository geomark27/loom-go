package helpers

import (
	"fmt"
	"net/http"
)

// AppError representa un error de aplicación con contexto
type AppError struct {
	Message    string
	StatusCode int
	Internal   error
}

func (e *AppError) Error() string {
	if e.Internal != nil {
		return fmt.Sprintf("%s: %v", e.Message, e.Internal)
	}
	return e.Message
}

// NewAppError crea un nuevo error de aplicación
func NewAppError(message string, statusCode int, internal error) *AppError {
	return &AppError{
		Message:    message,
		StatusCode: statusCode,
		Internal:   internal,
	}
}

// Wrap envuelve un error con contexto adicional
func Wrap(err error, message string) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %w", message, err)
}

// Unwrap desenvuelve un error
func Unwrap(err error) error {
	type unwrapper interface {
		Unwrap() error
	}

	u, ok := err.(unwrapper)
	if !ok {
		return nil
	}
	return u.Unwrap()
}

// Errores comunes predefinidos
var (
	ErrNotFound = &AppError{
		Message:    "Resource not found",
		StatusCode: http.StatusNotFound,
	}

	ErrBadRequest = &AppError{
		Message:    "Bad request",
		StatusCode: http.StatusBadRequest,
	}

	ErrUnauthorized = &AppError{
		Message:    "Unauthorized",
		StatusCode: http.StatusUnauthorized,
	}

	ErrForbidden = &AppError{
		Message:    "Forbidden",
		StatusCode: http.StatusForbidden,
	}

	ErrInternalServer = &AppError{
		Message:    "Internal server error",
		StatusCode: http.StatusInternalServerError,
	}

	ErrConflict = &AppError{
		Message:    "Resource conflict",
		StatusCode: http.StatusConflict,
	}

	ErrUnprocessable = &AppError{
		Message:    "Unprocessable entity",
		StatusCode: http.StatusUnprocessableEntity,
	}
)
