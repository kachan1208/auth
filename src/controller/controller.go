package controller

import (
	"github.com/kachan1208/auth/src/api"
	"github.com/kachan1208/auth/src/dao"
	"github.com/kachan1208/auth/src/model"
)

type Controller struct {
	repo *dao.TokenRepo
}

func NewController(repo *dao.TokenRepo) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) AuthByToken(req *api.AuthByTokenReq) (*model.Token, error) {
	t, err := c.repo.GetTokenByToken(req.Token)
	if err != nil {
		return nil, err
	}

	if t.IsEnabled != true {
		return nil, api.ErrTokenIsDisabled
	}

	return t, nil
}

func (c *Controller) TokenList(req *api.TokenListReq) ([]model.Token, error) {
	return c.repo.TokenList(req.AccountID)
}

func (c *Controller) GetToken(req *api.GetTokenReq) (*model.Token, error) {
	return c.repo.GetTokenByID(req.ID)
}

func (c *Controller) UpdateToken(req *api.UpdateTokenReq) error {
	_, err := c.repo.GetTokenByID(req.ID)
	if err != nil {
		return err
	}
	return c.repo.UpdateToken(req.ID, req.AccountID, req.IsEnabled)
}

func (c *Controller) CreateToken(req *api.CreateTokenReq) (*model.Token, error) {
	return c.repo.CreateToken(req.AccountID)
}

func (c *Controller) DeleteToken(req *api.DeleteTokenReq) error {
	return c.repo.DeleteToken(req.ID, req.AccountID)
}
