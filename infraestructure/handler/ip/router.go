package ip

import (
	"github.com/XelaMP/getIP.git/model"
	"github.com/labstack/echo/v4"
)

func NewRouter(specification model.RouterSpecification) {
	handler := buildHandler()

	publicRoutes(specification.Api, handler)
}

func buildHandler() handler {
	return newHandler()
}

func publicRoutes(api *echo.Echo, h handler, middlewares ...echo.MiddlewareFunc) {
	route := api.Group("v1/ip-client", middlewares...)

	route.GET("", h.ipClient)
}
