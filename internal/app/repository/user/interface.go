package user

import (
	"context"
)

type Repository interface {
	CreateUser(ctx context.Context, param User) (returnData Res, err error)
	GetUserByField(ctx context.Context, field string, value string) (returnData Res, err error)
}
