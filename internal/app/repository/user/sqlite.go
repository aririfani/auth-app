package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aririfani/auth-app/config"
	"time"
)

type db struct {
	Cfg               config.Config
	DB                *sql.DB
	StmtGetByUsername *sql.Stmt
}

func NewDB(cfg config.Config, d *sql.DB) Repository {
	return &db{
		Cfg: cfg,
		DB:  d,
	}
}

func (d *db) CreateUser(ctx context.Context, param User) (returnData CreateRes, err error) {
	stmt, err := d.DB.Prepare(QueryCreateUser)

	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	defer stmt.Close()

	result, err := stmt.Exec(
		param.UserName,
		param.Phone,
		time.Now(),
		param.Password,
		param.Role,
	)

	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	lastInsertID, err := result.LastInsertId()

	if err != nil {
		return
	}

	returnData = CreateRes{
		ID:       lastInsertID,
		UserName: param.UserName,
		Phone:    param.Phone,
		Role:     param.Role,
	}

	return
}

func (d *db) GetUserByUsername(ctx context.Context, userName string) (returnData Res, err error) {
	if d.StmtGetByUsername == nil {
		d.StmtGetByUsername, err = d.DB.Prepare(QueryGetUserByUsername)
		if err != nil {
			err = errors.New(fmt.Sprintf("%s", err))
			return
		}
	}

	rows, err := d.StmtGetByUsername.QueryContext(ctx, userName)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	defer func() {
		_ = rows.Close()
	}()

	count := 0
	for rows.Next() {
		count += 1
		err = rows.Scan(
			&returnData.ID,
			&returnData.UserName,
			&returnData.Phone,
			&returnData.Password,
			&returnData.RegisteredAt,
			&returnData.Role,
		)

		if err != nil {
			err = errors.New(fmt.Sprintf("%s", err))
			return
		}
	}

	if count == 0 {
		err = sql.ErrNoRows
		return
	}

	return
}
