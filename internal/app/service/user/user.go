package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/app/repository"
	"github.com/aririfani/auth-app/internal/app/repository/user"
	"github.com/aririfani/auth-app/internal/pkg/token"
	helper "github.com/aririfani/auth-app/internal/pkg/utils/generalhelper"
	"github.com/ulule/deepcopier"
	"time"
)

type srv struct {
	Repo   repository.Repositories
	Token  token.IToken
	Config config.Config
}

func NewSrv(repo repository.Repositories, token token.IToken, config config.Config) (returnData Service) {
	return &srv{
		Repo:   repo,
		Token:  token,
		Config: config,
	}
}

func (s *srv) CreateUser(ctx context.Context, u User) (returnData user.Res, err error) {
	returnData, err = s.Repo.User().GetUserByField(ctx, "username", u.UserName)

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
	userRepo, err := s.Repo.User().GetUserByField(ctx, "username", userName)
	_ = deepcopier.Copy(userRepo).To(&returnData)

	return
}

func (s *srv) GetUserByPhone(ctx context.Context, phone string) (returnData user.Res, err error) {
	userRepo, err := s.Repo.User().GetUserByField(ctx, "phone", phone)
	_ = deepcopier.Copy(userRepo).To(&returnData)
	return
}

func (s *srv) Login(ctx context.Context, login UserLogin) (returnData UserLoginRes, err error) {
	var userRepo user.Res
	userRepo, err = s.Repo.User().GetUserByField(ctx, "phone", login.Phone)
	if err == sql.ErrNoRows {
		err = errors.New("phone number invalid")
		return
	}

	validPassword, err := helper.CompareStringValid(login.Password, userRepo.Password)
	if !validPassword {
		err = errors.New("phone or password invalid")
		return
	}

	var getTokenRes token.GetToken
	var accessTokenExpiresAt int64
	var expiredAt uint64 = uint64(s.Config.GetInt("app.token_expired_at"))

	getTokenRes.AccessToken, accessTokenExpiresAt, err = s.Token.GetImplicitToken(
		token.Payload{
			Username:     userRepo.UserName,
			Phone:        userRepo.Phone,
			Role:         userRepo.Role,
			RegisteredAt: userRepo.RegisteredAt,
		}, expiredAt)

	getTokenRes.AccessTokenExpiresAt = time.Unix(accessTokenExpiresAt, 0)

	returnData = UserLoginRes{
		AccessToken: getTokenRes.AccessToken,
		ExpiredAt:   getTokenRes.AccessTokenExpiresAt,
	}

	return
}
