package service

import (
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/app/repository"
	"github.com/aririfani/auth-app/internal/app/service/user"
	"github.com/aririfani/auth-app/internal/pkg/token"
)

type Services interface {
	User() user.Service
}

type service struct {
	user user.Service
}

func NewService(repo repository.Repositories, config config.Config) Services {
	tokenPkg := token.New(
		token.WithIssuer(config.GetString("app.issuer")),
		token.WithSecretKey(config.GetString("app.secret_key")),
	)

	u := new(service)
	u.user = user.NewSrv(repo, tokenPkg, config)
	return u
}

func (u *service) User() user.Service {
	return u.user
}
