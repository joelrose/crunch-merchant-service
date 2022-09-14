package utils

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/labstack/gommon/log"
	"golang.org/x/exp/slices"
)

func ConvertOrderItemsToDto(orderItems []models.OrderItem) []dtos.OrderItem {
	var parentOrders []dtos.OrderItem

	for _, orderItem := range orderItems {
		if !orderItem.ParentId.Valid {
			parentOrders = append(parentOrders, ConvertOrderItemToDto(orderItem))
		}
	}

	for _, orderItem := range orderItems {
		if orderItem.ParentId.Valid {
			ind := slices.IndexFunc(parentOrders, func(c dtos.OrderItem) bool {
				return orderItem.ParentId.Int32 == int32(c.Id)
			})

			if ind != -1 {
				parentOrders[ind].SubItems = append(parentOrders[ind].SubItems, ConvertOrderItemToDto(orderItem))
			} else {
				// TODO: allow complexer items
				log.Errorf("... Parent not found for order item %v", orderItem)
			}
		}
	}

	return parentOrders
}

func ConvertOrderItemToDto(orderItem models.OrderItem) dtos.OrderItem {
	return dtos.OrderItem{
		Id:       orderItem.Id,
		Plu:      orderItem.Plu,
		Name:     orderItem.Name,
		Price:    orderItem.Price,
		Quantity: orderItem.Quantity,
		SubItems: []dtos.OrderItem{},
	}
}
