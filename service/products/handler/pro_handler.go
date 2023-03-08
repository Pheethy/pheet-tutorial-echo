package handler

import (
	"net/http"
	"pheet-tutorial-echo/service/products"

	"github.com/labstack/echo/v4"
)

type productHandler struct {
	proUs products.ProductUsecase
}

func NewProductHandler(proUs products.ProductUsecase) products.ProductHandler {
	return productHandler{proUs: proUs}
}

func (p productHandler) FetchProducts(c echo.Context) error {
	var ctx = c.Request().Context()

	products, err := p.proUs.GetProducts(ctx); 
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	var resp = map[string]interface{}{
		"products": products,
	}

	return c.JSON(http.StatusOK, resp)
}