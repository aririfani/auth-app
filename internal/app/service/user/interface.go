package user

import (
	"context"
	"github.com/aririfani/auth-app/internal/app/repository/user"
)

type Service interface {
	CreateUser(ctx context.Context, user User) (returnData user.CreateRes, err error)
	GetUserByUsername(ctx context.Context, userName string) (returnData user.Res, err error)
}
