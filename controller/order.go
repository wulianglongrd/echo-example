package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type OrderController interface {
	CreateOrder(c echo.Context) error
	GetPayInfo(c echo.Context) error
}

type orderController struct {
}

func NewOrderController() OrderController {
	return &orderController{}
}

// CreateOrder http://localhost:1323/createOrder
func (controller *orderController) CreateOrder(c echo.Context) error {
	return c.String(http.StatusOK, "create success")
}

// GetPayInfo http://localhost:1323/getPayInfo
func (controller *orderController) GetPayInfo(c echo.Context) error {
	order := struct {
		Title    string `json:"title"`
		TotalFee int    `json:"total_fee"`
	}{
		Title:    "book",
		TotalFee: 10,
	}
	return c.JSON(http.StatusOK, order)
}
