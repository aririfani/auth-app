package repository

import (
	"github.com/aririfani/auth-app/fetch-service/config"
	"github.com/aririfani/auth-app/fetch-service/internal/app/repository/commodity"
	"github.com/aririfani/auth-app/fetch-service/internal/pkg/currency"
	"net/http"
)

type Repositories interface {
	Commodity() commodity.Repository
}

type repository struct {
	commodity commodity.Repository
}

func NewRepo(cfg config.Config, httpClient *http.Client) Repositories {
	currencyPkg := currency.NewCurrency(httpClient)

	repo := new(repository)
	repo.commodity = commodity.NewRepo(commodity.NewHttpClient(cfg, httpClient, currencyPkg))

	return repo
}

func (r *repository) Commodity() commodity.Repository {
	return r.commodity
}
