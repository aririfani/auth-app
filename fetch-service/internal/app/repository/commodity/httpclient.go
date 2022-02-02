package commodity

import (
	"encoding/json"
	"github.com/aririfani/auth-app/fetch-service/config"
	"github.com/aririfani/auth-app/fetch-service/internal/pkg/currency"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

type httpClient struct {
	netClient *http.Client
	cfg       config.Config
	request   *http.Request
	response  *http.Response
	currency  currency.ICurrency
}

func NewHttpClient(cfg config.Config, nc *http.Client, iCurrency currency.ICurrency) Repository {
	return &httpClient{
		netClient: nc,
		cfg:       cfg,
		currency:  iCurrency,
	}
}

func (h *httpClient) FetchDataCommodity() (res []Commodity, err error) {
	h.request, err = http.NewRequest("GET", h.cfg.GetString("clients.efishery.commodity_uri"), nil)
	if err != nil {
		return
	}

	h.response, err = h.netClient.Do(h.request)
	if err != nil {
		return
	}

	defer h.response.Body.Close()

	var tmp []Commodity
	err = json.NewDecoder(h.response.Body).Decode(&tmp)
	if err != nil {
		return
	}

	idr, err := h.currency.GetConversionRateIDRtoUSD()
	if err != nil {
		return
	}

	for _, val := range tmp {
		priceIDR, _ := strconv.ParseFloat(val.Price, 8)
		priceUsd := priceIDR / idr.ConversionRate

		response := Commodity{
			Komoditas:    val.Komoditas,
			AreaProvinsi: val.AreaProvinsi,
			AreaKota:     val.AreaKota,
			Size:         val.AreaKota,
			Price:        val.Price,
			PriceUSD:     priceUsd,
			TglParsed:    val.TglParsed,
			Uuid:         val.Uuid,
		}

		res = append(res, response)
	}

	return
}

func (h *httpClient) FetchDataCommodityByProvince(provinceName string, sorted bool, sortedType string) (res []Commodity, err error) {
	h.request, err = http.NewRequest("GET", h.cfg.GetString("clients.efishery.commodity_uri"), nil)
	if err != nil {
		return
	}

	h.response, err = h.netClient.Do(h.request)
	if err != nil {
		return
	}

	defer h.response.Body.Close()

	var tmp []Commodity
	err = json.NewDecoder(h.response.Body).Decode(&tmp)
	if err != nil {
		return
	}

	idr, err := h.currency.GetConversionRateIDRtoUSD()
	if err != nil {
		return
	}

	for _, val := range tmp {
		priceIDR, _ := strconv.ParseFloat(val.Price, 8)
		priceUsd := priceIDR / idr.ConversionRate

		if val.AreaProvinsi == strings.ToUpper(provinceName) {
			response := Commodity{
				Komoditas:    val.Komoditas,
				AreaProvinsi: val.AreaProvinsi,
				AreaKota:     val.AreaKota,
				Size:         val.AreaKota,
				Price:        val.Price,
				PriceUSD:     priceUsd,
				TglParsed:    val.TglParsed,
				Uuid:         val.Uuid,
			}
			res = append(res, response)
		}

		if res != nil && sorted == true {
			var lowToHigh = false
			if sortedType == "low_to_high" {
				lowToHigh = true
			}

			SortCommodityByPrice(res, lowToHigh)
		}
	}

	return res, nil
}

func SortCommodityByPrice(commodity []Commodity, lowToHigh bool) {
	sort.Slice(commodity, func(i, j int) bool {
		if lowToHigh {
			return commodity[i].PriceUSD < commodity[j].PriceUSD
		} else {
			return commodity[i].PriceUSD > commodity[j].PriceUSD

		}
	})
}
