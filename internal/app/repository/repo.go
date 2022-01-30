package repository

import (
	"errors"
	"fmt"
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/app/driver/db"
	"github.com/aririfani/auth-app/internal/app/repository/user"
)

type Repositories interface {
	User() user.Repository
}

type repository struct {
	user user.Repository
}

func NewRepo(cfg config.Config) Repositories {
	dbase := db.New(db.WithConfig(cfg))
	sqlConn, err := dbase.Manager(db.SqlLiteDialectParam)

	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
	}

	repo := new(repository)
	repo.user = user.NewRepo(user.NewDB(cfg, sqlConn))

	return repo
}

func (r *repository) User() user.Repository {
	return r.user
}
