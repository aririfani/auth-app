package user

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, param User) (returnData Res, err error)
	GetUserByUsername(ctx context.Context, userName string) (returnData Res, err error)
}
