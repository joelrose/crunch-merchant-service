package utils

import (
	"github.com/joelrose/crunch-merchant-service/dtos"
)

func CalculateOrderPrice(items []dtos.OrderItem) int {
	if len(items) == 0 {
		return 0
	}

	var price int
	for _, item := range items {
		if item.SubItems != nil {
			price += item.Quantity * (CalculateOrderPrice(item.SubItems) + item.Price)
		} else {
			price += item.Quantity * item.Price
		}
	}

	return price
}
