package currency

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ICurrency interface {
	GetConversionRateIDRtoUSD() (res Currency, err error)
}

type currency struct {
	httpClient *http.Client
	request    *http.Request
	response   *http.Response
}

func NewCurrency(client *http.Client) ICurrency {
	return &currency{
		httpClient: client,
	}
}

func (c *currency) GetConversionRateIDRtoUSD() (res Currency, err error) {
	c.request, err = http.NewRequest("GET", "https://v6.exchangerate-api.com/v6/a4e05e550f391fce03c45f60/pair/USD/IDR", nil)
	if err != nil {
		return
	}

	c.response, err = c.httpClient.Do(c.request)
	if err != nil {
		return
	}

	defer c.response.Body.Close()
	err = json.NewDecoder(c.response.Body).Decode(&res)

	if err != nil {
		fmt.Println(err)
		return
	}

	return
}
