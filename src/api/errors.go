package api

import (
	"net/http"

	"github.com/kachan1208/auth/src/transport/http/errors"
)

var (
	ErrInvalidJSON = errors.NewHTTPError(
		http.StatusBadRequest,
		4001,
		"invalid json",
	)

	ErrInvalidBase64 = errors.NewHTTPError(
		http.StatusBadRequest,
		4002,
		"invalid base64",
	)

	ErrAccountIDIsNotSet = errors.NewHTTPError(
		http.StatusBadRequest,
		4003,
		"'Account-Id' header is not set",
	)

	ErrTokenIsNotSet = errors.NewHTTPError(
		http.StatusBadRequest,
		4004,
		"'Authorization' header is not set",
	)

	ErrNotFound = errors.NewHTTPError(
		http.StatusNotFound,
		4040,
		"token not found",
	)

	ErrTokenIDIsInvalid = errors.NewHTTPError(
		http.StatusBadRequest,
		4005,
		"token id is invalid",
	)

	ErrAccountIDMistmatch = errors.NewHTTPError(
		http.StatusBadRequest,
		4006,
		"token can't be removed, invalid 'Account-Id' header",
	)
)
