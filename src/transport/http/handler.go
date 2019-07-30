package http

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-kit/kit/log"
	transport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"

	"github.com/kachan1208/auth/src/api"
	"github.com/kachan1208/auth/src/controller"
)

const (
	RouteAuthByToken = "/token/auth"
	RouteCreateToken = "/token"
	RouteDeleteToken = "/token/{id}"
	RouteHealth      = "/health"
)

type Handler struct {
	Address string
	Router  *mux.Router
}

func NewHandler(address string, logger log.Logger, c *controller.Controller) *Handler {
	options := []transport.ServerOption{
		transport.ServerErrorLogger(logger),
		transport.ServerErrorEncoder(handleError),
	}

	authByToken := transport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return c.AuthByToken(request.(*api.AuthByTokenReq))
		},
		unmarshalAuthByTokenReq,
		marshalAuthByTokenResp,
		options...,
	)
	createToken := transport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return c.CreateToken(request.(*api.CreateTokenReq))
		},
		unmarshalCreateTokenReq,
		marshalCreateTokenResp,
		options...,
	)
	deleteToken := transport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return nil, c.DeleteToken(request.(*api.DeleteTokenReq))
		},
		unmarshalDeleteTokenReq,
		marshalDeleteTokenResp,
		options...,
	)

	health := transport.NewServer(
		func(ctx context.Context, request interface{}) (interface{}, error) {
			return nil, nil
		},
		func(_ context.Context, r *http.Request) (request interface{}, err error) {
			return nil, nil
		},
		func(ctx context.Context, w http.ResponseWriter, response interface{}) error {
			w.WriteHeader(http.StatusOK)
			return nil
		},
		options...,
	)

	router := mux.NewRouter()
	router.Methods("POST").Path(RouteAuthByToken).Handler(authByToken)
	router.Methods("POST").Path(RouteCreateToken).Handler(createToken)
	router.Methods("DELETE").Path(RouteDeleteToken).Handler(deleteToken)
	router.Methods("GET").Path(RouteHealth).Handler(health)

	return &Handler{
		Address: address,
		Router:  router,
	}
}

func unmarshalAuthByTokenReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := api.AuthByTokenReq{}
	req.Token = r.Header.Get("Authorization")
	if req.Token == "" || len(req.Token) != 64 {
		return nil, api.ErrTokenIsNotSet
	}

	return &req, err
}

func marshalAuthByTokenResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok {
		handleError(ctx, err, w)
		return err
	}

	return nil
}

func unmarshalCreateTokenReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := api.CreateTokenReq{}
	req.AccountID = r.Header.Get("Account-Id")
	if req.AccountID == "" || len(req.AccountID) != 32 {
		return nil, api.ErrAccountIDIsNotSet
	}

	return &req, nil
}

func marshalCreateTokenResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok {
		handleError(ctx, err, w)
		return err
	}
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

func unmarshalDeleteTokenReq(_ context.Context, r *http.Request) (request interface{}, err error) {
	req := api.DeleteTokenReq{}
	req.AccountID = r.Header.Get("Account-Id")
	if req.AccountID == "" || len(req.AccountID) != 32 {
		return nil, api.ErrAccountIDIsNotSet
	}

	req.ID = mux.Vars(r)["id"]
	if len(req.ID) != 32 {
		return nil, api.ErrTokenIDIsInvalid
	}

	return &req, nil
}

func marshalDeleteTokenResp(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok {
		handleError(ctx, err, w)
		return err
	}

	w.WriteHeader(http.StatusOK)
	return nil
}
