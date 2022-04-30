package router

import (
	"github.com/labstack/echo/v4"
	"github.com/wulianglongrd/echo-example/controller"
	"net/http"
)

func SetupRouter(e *echo.Echo) {

	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })

	order := controller.NewOrderController()
	e.GET("/createOrder", func(c echo.Context) error { return order.CreateOrder(c) })
	e.GET("/getPayInfo", func(c echo.Context) error { return order.GetPayInfo(c) })

	bill := controller.NewBillController()
	e.GET("/changeBill", func(c echo.Context) error { return bill.ChangeBill(c) })
	e.GET("/renderBill", func(c echo.Context) error { return bill.RenderBill(c) })
}
