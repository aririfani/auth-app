package service

import (
	"github.com/aririfani/auth-app/internal/app/repository"
	"github.com/aririfani/auth-app/internal/app/service/user"
)

type Services interface {
	User() user.Service
}

type service struct {
	user user.Service
}

func NewService(repo repository.Repositories) Services {
	u := new(service)
	u.user = user.NewSrv(repo)
	return u
}

func (u *service) User() user.Service {
	return u.user
}
