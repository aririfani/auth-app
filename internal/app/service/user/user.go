package user

import (
	"context"
	"github.com/aririfani/auth-app/internal/app/repository"
	"github.com/aririfani/auth-app/internal/app/repository/user"
	"github.com/ulule/deepcopier"
)

type srv struct {
	Repo repository.Repositories
}

func NewSrv(repo repository.Repositories) (returnData Service) {
	return &srv{
		Repo: repo,
	}
}

func (s *srv) CreateUser(ctx context.Context, u User) (returnData user.CreateRes, err error) {
	var userRepo user.User
	_ = deepcopier.Copy(u).To(&userRepo)
	returnData, err = s.Repo.User().CreateUser(ctx, userRepo)

	return
}

func (s *srv) GetUserByUsername(ctx context.Context, userName string) (returnData user.Res, err error) {
	userRepo, err := s.Repo.User().GetUserByUsername(ctx, userName)
	_ = deepcopier.Copy(userRepo).To(&returnData)

	return
}
