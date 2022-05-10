package app

import (
	"github.com/raflynagachi/go-store/models/order"
	"github.com/raflynagachi/go-store/models/product"
	"github.com/raflynagachi/go-store/models/user"
)

type Model struct {
	Model interface{}
}

func RegisterModels() []Model {
	return []Model{
		{Model: user.User{}},
		{Model: user.Address{}},
		{Model: product.Product{}},
		{Model: product.ProductImage{}},
		{Model: product.Section{}},
		{Model: product.Category{}},
		{Model: order.Order{}},
		{Model: order.OrderItem{}},
		{Model: order.OrderCustomer{}},
		{Model: order.Payment{}},
		{Model: order.Shipment{}},
	}
}
