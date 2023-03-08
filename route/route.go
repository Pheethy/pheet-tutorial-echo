package route

import (
	"pheet-tutorial-echo/service/products"

	"github.com/labstack/echo/v4"
)

type Route struct {
	e     *echo.Echo
}

func NewRoute(e *echo.Echo) *Route {
	return &Route{e: e}
}

func (r Route) RegisterProduct(handler products.ProductHandler) {
	r.e.GET("/v1/products", handler.FetchProducts)
}