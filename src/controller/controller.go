package controller

import (
	"context"

	"github.com/kachan1208/auth/src/dao"
)

type Controller struct {
	repo *dao.TokenRepository
}

func NewController(repo *dao.TokenRepository) *Controller {
	return &Controller{
		repo: repo,
	}
}

func (c *Controller) AuthByToken(ctx context.Context, req *api.AuthByTokenReq) error {
	token, err := c.db.GetToken(req.Token)

}
