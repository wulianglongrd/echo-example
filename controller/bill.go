package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/wulianglongrd/echo-example/model"
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
	bill := model.Bill{
		Name:  "基础费",
		Type:  "basic_fee",
		Value: 100,
	}
	return c.JSON(http.StatusOK, bill)
}
