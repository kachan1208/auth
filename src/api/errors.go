package api

import (
	"net/http"

	"github.com/kachan1208/transport/http/errors"
)

var (
	ErrInvalidJSON := NewHTTPError(
		http.StatusBadRequest,
		4001,
		"invalid json",
	)

	ErrInvalidBase64 := NewHTTPError(
		http.StatusBadRequest,
		4002,
		"invalid base64",
	)

	ErrAccountIDIsNotSet := NewHTTPError(
		http.StatusBadRequest,
		4003,
		"'Account-Id' header is not set",
	)

	ErrTokenIsNotSet := NewHTTPError(
		http.StatusBadRequest,
		4004,
		"'Authorization' header is not set",
	)

	ErrNotFound := NewHTTPError(
		http.StatusNotFound,
		4040,
		"token not found",
	)
)