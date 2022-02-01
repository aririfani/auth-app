package commodity

import (
	"github.com/aririfani/auth-app/fetch-service/internal/app/service"
	"net/http"
)

type Handler interface {
	GetCommodity(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
}

type commodityServ struct {
	service service.Services
}

func NewCommodity(srv service.Services) Handler {
	return &commodityServ{
		service: srv,
	}
}

func (u *commodityServ) GetCommodity(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	resp, err = u.service.Commodity().FetchDataCommodity()
	w.Header().Add("Content-Type", "application/json")

	return
}
