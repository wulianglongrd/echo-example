package controller

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type BillController interface {
	ChangeBill(c echo.Context) error
	RenderBill(c echo.Context) error
}

type billController struct {
}

func NewBillController() BillController {
	return &billController{}
}

func (controller *billController) ChangeBill(c echo.Context) error {
	return c.String(http.StatusOK, "change success")
}

func (controller *billController) RenderBill(c echo.Context) error {
	bill := struct {
		FeeName string `json:"fee_name"`
		FeeType string `json:"fee_type"`
		Fee     int    `json:"fee"`
	}{
		FeeName: "基础费",
		FeeType: "basic_fee",
		Fee:     100,
	}
	return c.JSON(http.StatusOK, bill)
}
