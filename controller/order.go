package controller

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/wulianglongrd/echo-example/service"
	"github.com/wulianglongrd/echo-example/service/vo"
	"net/http"
)

type OrderController interface {
	CreateOrder(c echo.Context) error
	GetPayInfo(c echo.Context) error
}

type orderController struct {
	service service.OrderService
}

func NewOrderController() OrderController {
	return &orderController{service: service.NewOrderService()}
}

// CreateOrder http://localhost:1323/createOrder?buyer_id=123456&order_id=123&title=book&total_fee=100
func (ctl *orderController) CreateOrder(c echo.Context) error {
	orderVo := &vo.CreateOrderVo{}
	err := c.Bind(orderVo)
	if err != nil {
		fmt.Println("bind orderVo failed")
		return err
	}

	_, err = ctl.service.CreateOrder(orderVo)
	if err != nil {
		fmt.Println("createOrder failed", err)
		return err
	}

	return c.String(http.StatusOK, "create success")
}

// GetPayInfo http://localhost:1323/getPayInfo?buyer_id=123456&order_id=123
func (ctl *orderController) GetPayInfo(c echo.Context) error {
	buyerId := c.QueryParam("buyer_id")
	orderId := c.QueryParam("order_id")

	order, err := ctl.service.FindOrder(buyerId, orderId)
	if err != nil {
		fmt.Println("findOrder failed", err)
		return err
	}
	return c.JSON(http.StatusOK, order)
}
