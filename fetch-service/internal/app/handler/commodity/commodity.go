package commodity

import (
	"errors"
	"github.com/aririfani/auth-app/fetch-service/internal/app/service"
	"net/http"
	"strconv"
)

type Handler interface {
	GetCommodity(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
	GetDataCommodityByProvince(w http.ResponseWriter, r *http.Request) (resp interface{}, err error)
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

func (u *commodityServ) GetDataCommodityByProvince(w http.ResponseWriter, r *http.Request) (resp interface{}, err error) {
	provinceName := r.URL.Query().Get("province_name")
	sorted, err := strconv.ParseBool(r.URL.Query().Get("is_sorted"))

	if err != nil {
		err = errors.New("query parameter is_sorted required")
		return
	}

	sortedType := r.URL.Query().Get("sorted_type")

	resp, err = u.service.Commodity().FetchDataCommodityByProvince(provinceName, sorted, sortedType)
	w.Header().Add("Content-Type", "application/json")

	return
}
