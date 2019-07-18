package controller

import (
	"context"

	"github.com/kachan1208/auth/src/dao"
	"github.com/kachan1208/auth/src/api"
)

type Controller struct {
	repo *dao.TokenRepo
}

func NewController(repo *dao.TokenRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) AuthByToken(ctx context.Context, req *api.AuthByTokenReq) (*model.Token, error) {
	token, err := c.repo.GetToken(req.Token)
	
	return token, err
}

func (c *Controller) CreateToken(ctx context.Context, req *api.CreateTokenReq) (*model.Token, error) {
	token, err := c.repo.CreateToken(req.AccountID)

	return token, err
}


func (c *Controller) DeleteToken(ctx context.Context, req *api.CreateTokenReq) error {
	token, err := c.repo.GetToken(req.ID)
	if err != nil {
		return err
	}

	if token.AccountID != req.AccountID {
		return api.ErrAccountIDMistmatch
	}

	return c.repo.DeleteToken(req.ID)
}