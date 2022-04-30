package service

import (
	"fmt"
	"github.com/wulianglongrd/echo-example/model"
	"github.com/wulianglongrd/echo-example/service/vo"
)

type OrderService interface {
	CreateOrder(vo *vo.CreateOrderVo) (bool, error)
	FindOrder(buyerId string, orderId string) (model.Order, error)
}

type orderService struct {
}

func NewOrderService() OrderService {
	return &orderService{}
}

func (service orderService) CreateOrder(vo *vo.CreateOrderVo) (bool, error) {
	fmt.Printf("CreateOrder params: %+v\n", vo)
	return true, nil
}

func (service orderService) FindOrder(buyerId string, orderId string) (model.Order, error) {
	fmt.Printf("FindOrder buyerId: %s, orderId: %s\n", buyerId, orderId)

	order := model.Order{
		ID:       1,
		BuyerId:  buyerId,
		OrderId:  orderId,
		Title:    "book order",
		TotalFee: 100,
	}
	return order, nil
}
