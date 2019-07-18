package api

type AuthByTokenReq struct {
	Token string
}

type AuthByTokenResp struct {
	AccountID string 
}

type CreateTokenReq struct {
	AccountID string
}

type CreateTokenResp struct {
	Token string `json:"token"`
}

type DeleteTokenReq struct {
	ID string
	AccountID string
}