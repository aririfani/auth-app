package commodity

import (
	"encoding/json"
	"fmt"
	"github.com/patrickmn/go-cache"
	"strconv"
	"time"
)

type repo struct {
	DB         Repository
	httpClient *httpClient
	cache      *cache.Cache
}

func NewRepo(db Repository, c *cache.Cache) (returnData Repository) {
	return &repo{
		DB:    db,
		cache: c,
	}
}

func (r *repo) FetchDataCommodity() (returnData []Commodity, err error) {
	rd := new([]Commodity)
	data, found := r.cache.Get("FetchDataCommodity")

	if found {
		d, _ := json.Marshal(data)
		err = json.Unmarshal(d, &rd)
		if err != nil {
			return
		}

		returnData = *rd
		return returnData, nil
	}

	returnData, err = r.DB.FetchDataCommodity()
	if err != nil {
		return
	}

	r.cache.Set("FetchDataCommodity", returnData, 2*time.Minute)

	return
}

func (r *repo) FetchDataCommodityByProvince(provinceName string, sorted bool, sortedType string) (returnData []Commodity, err error) {
	rd := new([]Commodity)
	data, found := r.cache.Get(fmt.Sprintf("FetchDataCommodityByProvince:%s%s%s", provinceName, strconv.FormatBool(sorted), sortedType))
	if found {
		d, _ := json.Marshal(data)
		err = json.Unmarshal(d, &rd)
		if err != nil {
			return
		}

		returnData = *rd
		return returnData, nil
	}

	returnData, err = r.DB.FetchDataCommodityByProvince(provinceName, sorted, sortedType)
	if err != nil {
		return
	}

	if returnData != nil {
		r.cache.Set(fmt.Sprintf("FetchDataCommodityByProvince:%s%s%s", provinceName, strconv.FormatBool(sorted), sortedType), returnData, 2*time.Minute)
	}

	return
}
