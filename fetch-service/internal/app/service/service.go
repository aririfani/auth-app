package service

import (
	"github.com/aririfani/auth-app/fetch-service/config"
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository"
	"github.com/aririfani/auth-app/fetch-service/internal/app/service/commodity"
)

type Services interface {
	Commodity() commodity.Service
}

type service struct {
	commodity commodity.Service
}

func NewService(repo repository.Repositories, config config.Config) Services {
	u := new(service)
	u.commodity = commodity.NewSrv(repo)
	return u
}

func (u *service) Commodity() commodity.Service {
	return u.commodity
}
