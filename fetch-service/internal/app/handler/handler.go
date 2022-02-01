package handler

import (
	"github.com/aririfani/auth-app/fetch-service/config"
	"github.com/aririfani/auth-app/fetch-service/internal/app/handler/commodity"
	"github.com/aririfani/auth-app/fetch-service/internal/app/service"
)

type Handler interface {
	CommodityHandler() commodity.Handler
}

type handler struct {
	commodity commodity.Handler
}

func NewHandler(cfg config.Config, service service.Services) Handler {
	h := new(handler)
	h.commodity = commodity.NewCommodity(service)

	return h
}

func (h *handler) CommodityHandler() commodity.Handler {
	return h.commodity
}
