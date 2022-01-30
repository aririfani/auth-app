package handler

import (
	"github.com/aririfani/auth-app/config"
	"github.com/aririfani/auth-app/internal/app/handler/user"
	"github.com/aririfani/auth-app/internal/app/service"
)

type Handler interface {
	UserHandler() user.Handler
}

type handler struct {
	user user.Handler
}

func NewHandler(cfg config.Config, service service.Services) Handler {
	h := new(handler)
	h.user = user.NewHandler(service)

	return h
}

func (h *handler) UserHandler() user.Handler {
	return h.user
}
