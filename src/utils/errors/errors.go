package errors

import "net/http"

type ApiError interface {
	Status() int
	Message() string
	Error() string
}

type apiError struct {
	Astatus  int    `json:"status"`
	Amessage string `json:"message"`
	AnError  string `json:"error,omitempty"`
}

func (ae *apiError) Status() int {
	return ae.Astatus
}

func (ae *apiError) Message() string {
	return ae.Amessage
}

func (ae *apiError) Error() string {
	return ae.AnError
}

func NewNotFoundApiError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusNotFound,
		Amessage: message,
	}
}

func NewInternalServerError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusInternalServerError,
		Amessage: message,
	}
}

func NewBadRequestError(message string) ApiError {
	return &apiError{
		Astatus:  http.StatusBadRequest,
		Amessage: message,
	}
}

func NewApiError(statusCode int, message string) ApiError {
	return &apiError{
		Astatus:  statusCode,
		Amessage: message,
	}
}
