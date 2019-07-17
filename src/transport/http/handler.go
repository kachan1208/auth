package handler

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"

	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/kachan1208/auth/src/api"
)

const (
	RouteAuthByToken = "/token/auth"
	RouteCreateToken = "/token"
)

type Handler struct {
	address string
	router  *mux.Router
}

func NewHandler(address string) *Handler {
	options := []transport.ServerOptions{
		transport.ServerErrorEncoder(handleError),
	}

	authByToken := transport.NewServer(
		controller.AuthByToken,
		unmarshalAuthByTokenReq,
		marshalAuthByTokenResp,
	)
	createToken := transport.NewServer(
		controller.CreateToken,
		unmarshalCreateTokenReq,
		marshalCreateResp,
	)

	router := mux.NewRouter()
	router.Methods("Post").Path(RouteAuthByToken).Handler(authByToken)
	router.Methods("Post").Path(RouteCreateToken).Handler(createToken)

	return &Handler{
		address: address,
		router:  router,
	}
}

func unmarshalAuthByTokenReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := api.AuthByTokenReq{}
	req.Token = r.Header.Get["Authorization"]
	if req.Token == "" || len(req.Token) != 32 {
		return nil, api.ErrTokenIsNotSet
	}

	return &req, err
}

func marshalAuthByTokenResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := response.(error); err != nil {
		handleError(err)
		return err
	}

	return nil
}

func unmarshalCreateTokenReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := api.CreateTokenReq{}
	req.AccountID = r.Header.Get["Account-Id"]
	if req.AccountID == "" || len(req.AccountID) != 32 {
		return nil, api.ErrAccountIDIsNotSet
	}

	return &req, nil
}

func marshalCreateTokenResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err := response.(error); err != nil {
		handleError(err)
		return err
	}

	return json.NewEncoder(w).Encode(resp)
}

