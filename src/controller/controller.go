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
	return c.repo.GetToken(req.Token)
}

func (c *Controller) CreateToken(req *api.CreateTokenReq) (*model.Token, error) {
	return c.repo.CreateToken(req.AccountID)
}

func (c *Controller) DeleteToken(req *api.DeleteTokenReq) error {
	return c.repo.DeleteToken(req.ID, req.AccountID)
}
