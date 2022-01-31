package user

import (
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/aririfani/auth-app/internal/app/service"
	"github.com/aririfani/auth-app/internal/app/service/user"
	"github.com/golang-jwt/jwt"
	"io/ioutil"
	"net/http"
)

type Handler interface {
	CreateUser(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
	GetUserByUserName(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
	Login(w http.ResponseWriter, r *http.Request) (res interface{}, err error)
	GetClaimUser(w http.ResponseWriter, r *http.Request) (res interface{}, err error)
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

	w.Header().Add("Content-Type", "application/json")

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

	w.Header().Add("Content-Type", "application/json")

	return
}

func (u *userServ) Login(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	ctx := r.Context()
	bodyByte, _ := ioutil.ReadAll(r.Body)

	var request user.UserLogin
	err = json.Unmarshal(bodyByte, &request)
	if err != nil || request == (user.UserLogin{}) {
		err = errors.New("error payload")
		return
	}

	resp, err = u.service.User().Login(ctx, request)

	w.Header().Add("Content-Type", "application/json")

	return
}

func (u *userServ) GetClaimUser(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	resp = user.ClaimRes{
		UserName:     r.Context().Value("claims").(jwt.MapClaims)["Username"].(string),
		Phone:        r.Context().Value("claims").(jwt.MapClaims)["Phone"].(string),
		RegisteredAt: r.Context().Value("claims").(jwt.MapClaims)["RegisteredAt"].(string),
		Role:         r.Context().Value("claims").(jwt.MapClaims)["Role"].(string),
	}

	w.Header().Add("Content-Type", "application/json")

	return
}
