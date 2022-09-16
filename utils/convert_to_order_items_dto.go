package utils

import (
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
)

func convertOrderItemToDto(orderItem models.OrderItem) dtos.OrderItem {
	return dtos.OrderItem{
		Id:       orderItem.Id,
		Plu:      orderItem.Plu,
		Name:     orderItem.Name,
		Price:    orderItem.Price,
		Quantity: orderItem.Quantity,
		SubItems: []dtos.OrderItem{},
	}
}

func convertHelper(parentMap map[int][]models.OrderItem, orderItems []models.OrderItem) []dtos.OrderItem {
	var result []dtos.OrderItem
	for ind, orderItem := range orderItems {
		result = append(result, convertOrderItemToDto(orderItem))

		if len(parentMap[orderItem.Id]) > 0 {
			result[ind].SubItems = convertHelper(parentMap, parentMap[orderItem.Id])
		}
	}

	return result
}

func ConvertOrderItemsToDto(orderItems []models.OrderItem) []dtos.OrderItem {
	parentIdMap := make(map[int][]models.OrderItem)
	for ind, orderItem := range orderItems {
		if orderItem.ParentId.Valid {
			parentIdMap[int(orderItem.ParentId.Int32)] = append(parentIdMap[int(orderItem.ParentId.Int32)], orderItems[ind])
		}
	}

	result := []dtos.OrderItem{}
	for _, orderItem := range orderItems {
		if !orderItem.ParentId.Valid {
			res := convertOrderItemToDto(orderItem)

			if len(parentIdMap[orderItem.Id]) > 0 {
				res.SubItems = convertHelper(parentIdMap, parentIdMap[orderItem.Id])
			}

			result = append(result, res)
		}
	}

	return result
}
