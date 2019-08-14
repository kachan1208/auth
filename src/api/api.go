package api

import "github.com/kachan1208/auth/src/model"

type AuthByTokenReq struct {
	Token string
}

type AuthByTokenResp struct {
	AccountID string
}

type GetTokenReq struct {
	ID        string
	AccountID string
}

type GetTokenResp struct {
	Token string `json:"token"`
}

type CreateTokenReq struct {
	AccountID string
}

type CreateTokenResp struct {
	Token string `json:"token"`
}

type DeleteTokenReq struct {
	ID        string
	AccountID string
}
type UpdateTokenReq struct {
	ID        string
	AccountID string
	IsEnabled bool
}

type TokenListReq struct {
	AccountID string
}

type TokenListResp = model.Token
