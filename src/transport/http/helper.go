package http

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	"golang.org/x/xerrors"

	"github.com/kachan1208/auth/src/api"
	"github.com/kachan1208/auth/src/transport/http/errors"
)

func unmarshalJSON(reader io.Reader, obj interface{}) error {
	body, err := ioutil.ReadAll(reader)
	if err != nil {
		return api.ErrInvalidJSON
	}

	if err = json.Unmarshal(body, obj); err != nil {
		if _, ok := err.(base64.CorruptInputError); ok {
			return api.ErrInvalidBase64
		}
		return api.ErrInvalidJSON
	}

	return nil
}

func handleError(ctx context.Context, err error, w http.ResponseWriter) {
	var httpErr errors.HTTPError
	if err != nil && xerrors.As(err, &httpErr) {
		w.WriteHeader(httpErr.StatusCode)
		w.Write([]byte(err.Error()))
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"code":10000, "message": "internal server error"}`))
	}
}
