package user

import (
	"context"
	"github.com/aririfani/auth-app/internal/app/repository/user"
)

type Service interface {
	CreateUser(ctx context.Context, user User) (returnData user.Res, err error)
	GetUserByUsername(ctx context.Context, userName string) (returnData user.Res, err error)
	GetUserByPhone(ctx context.Context, phone string) (returnData user.Res, err error)
	Login(ctx context.Context, login UserLogin) (returnData UserLoginRes, err error)
}
