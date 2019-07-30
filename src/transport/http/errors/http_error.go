package errors

import "fmt"

type HTTPError struct {
	StatusCode int
	Code       int
	Message    string
}

func NewHTTPError(statusCode, code int, msg string) error {
	return HTTPError{
		StatusCode: statusCode,
		Code:       code,
		Message:    msg,
	}
}

func (e HTTPError) Error() string {
	return fmt.Sprintf(`{"code": %d, "message": %s}`, e.Code, e.Message)
}
