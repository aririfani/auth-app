package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/aririfani/auth-app/config"
	helper "github.com/aririfani/auth-app/internal/pkg/utils/generalhelper"
	"time"
)

type db struct {
	Cfg            config.Config
	DB             *sql.DB
	StmtGetByField *sql.Stmt
}

func NewDB(cfg config.Config, d *sql.DB) Repository {
	return &db{
		Cfg: cfg,
		DB:  d,
	}
}

func (d *db) CreateUser(ctx context.Context, param User) (returnData Res, err error) {
	stmt, err := d.DB.Prepare(QueryCreateUser)

	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	defer stmt.Close()

	generatePassword := helper.GenerateString(4)
	password, err := helper.EncryptString(generatePassword)

	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	result, err := stmt.Exec(
		param.UserName,
		param.Phone,
		time.Now(),
		password,
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

	returnData = Res{
		ID:       lastInsertID,
		UserName: param.UserName,
		Phone:    param.Phone,
		Role:     param.Role,
		Password: password,
	}

	return
}

func (d *db) GetUserByField(ctx context.Context, field string, value string) (returnData Res, err error) {
	if d.StmtGetByField != nil {
		d.StmtGetByField = nil
	}

	var query = fmt.Sprintf(QueryGetUserByField, field, "'"+value+"'")
	d.StmtGetByField, err = d.DB.Prepare(query)
	if err != nil {
		err = errors.New(fmt.Sprintf("%s", err))
		return
	}

	rows, err := d.StmtGetByField.QueryContext(ctx)
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
