package user

import "context"

type repo struct {
	DB Repository
}

func NewRepo(db Repository) (returnData Repository) {
	return &repo{
		DB: db,
	}
}

func (r *repo) CreateUser(ctx context.Context, param User) (returnData Res, err error) {
	returnData, err = r.DB.CreateUser(ctx, param)
	if err != nil {
		return
	}

	return
}

func (r *repo) GetUserByField(ctx context.Context, field string, value string) (returnData Res, err error) {
	returnData, err = r.DB.GetUserByField(ctx, field, value)
	if err != nil {
		return
	}

	return
}
