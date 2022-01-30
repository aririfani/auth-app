package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aririfani/auth-app/internal/app/repository"
	"github.com/aririfani/auth-app/internal/app/repository/user"
	helper "github.com/aririfani/auth-app/internal/pkg/utils/generalhelper"
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

func (s *srv) CreateUser(ctx context.Context, u User) (returnData user.Res, err error) {
	returnData, err = s.Repo.User().GetUserByUsername(ctx, u.UserName)

	if err == sql.ErrNoRows {
		var userRepo user.User
		_ = deepcopier.Copy(u).To(&userRepo)
		returnData, err = s.Repo.User().CreateUser(ctx, userRepo)
	}

	returnData.Password, err = helper.DecryptString(returnData.Password)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	return
}

func (s *srv) GetUserByUsername(ctx context.Context, userName string) (returnData user.Res, err error) {
	userRepo, err := s.Repo.User().GetUserByUsername(ctx, userName)
	_ = deepcopier.Copy(userRepo).To(&returnData)

	return
}
