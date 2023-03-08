package products

import (
	"github.com/labstack/echo/v4"
)

type ProductHandler interface {
	FetchProducts(c echo.Context) error
}