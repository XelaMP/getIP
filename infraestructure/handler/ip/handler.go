package ip

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type handler struct{}

func newHandler() handler {
	return handler{}
}

func (h handler) ipClient(c echo.Context) error {
	return c.JSON(http.StatusOK, c.RealIP())
}
