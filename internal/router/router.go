package router

import (
	"github.com/go-openapi/loads"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server"
	"github.com/mhsbz/xiaohan/api/gen/xiaohan/server/operations"
	"github.com/mhsbz/xiaohan/internal/services"
)

type Router struct {
	SwaggerAPI *operations.XiaohanAPI
	service    *services.Service
}

func NewRouter() *Router {
	swaggerSpec, err := loads.Analyzed(server.SwaggerJSON, "")
	if err != nil {
		panic(err)
	}
	return &Router{
		SwaggerAPI: operations.NewXiaohanAPI(swaggerSpec),
		service:    services.NewService(),
	}
}

func (router *Router) RegisterRoutes() {
	router.SwaggerAPI.ActionHandler = operations.ActionHandlerFunc(router.service.Action)
}
