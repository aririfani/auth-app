package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/aririfani/auth-app/internal/app/service"
	"github.com/aririfani/auth-app/internal/app/service/user"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
	GetUserByUserName(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
}

type userServ struct {
	service service.Services
}

func NewHandler(srv service.Services) Handler {
	return &userServ{
		service: srv,
	}
}

func (u *userServ) CreateUser(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request user.User
	err = json.Unmarshal(bodyByte, &request)
	if err != nil || request == (user.User{}) {
		err = errors.New("error payload")
		return
	}

	resp, err = u.service.User().CreateUser(ctx, request)

	return
}

func (u *userServ) GetUserByUserName(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	ctx := r.Context()
	userName := r.URL.Query().Get("username")

	resp, err = u.service.User().GetUserByUsername(ctx, userName)

	if err == sql.ErrNoRows {
		err = errors.New("err data notfound")
		return
	}

	return
}
