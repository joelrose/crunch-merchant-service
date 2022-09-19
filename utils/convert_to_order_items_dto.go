package utils

import (
	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
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

func convertHelper(parentMap map[uuid.UUID][]models.OrderItem, orderItems []models.OrderItem) []dtos.OrderItem {
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
	parentIdMap := make(map[uuid.UUID][]models.OrderItem)
	for ind, orderItem := range orderItems {
		if orderItem.ParentId.Valid {
			parentIdMap[orderItem.ParentId.UUID] = append(parentIdMap[orderItem.ParentId.UUID], orderItems[ind])
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
