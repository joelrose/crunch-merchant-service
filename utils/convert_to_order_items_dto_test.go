package utils

import (
	"testing"

	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/stretchr/testify/assert"
)

func TestOneSimpleItem(t *testing.T) {
	input := []models.OrderItem{
		{
			Id:       1,
			Plu:      "123",
			Name:     "test",
			Price:    100,
			Quantity: 1,
			OrderId:  1,
		},
	}

	expected := []dtos.OrderItem{
		{
			Plu:      "123",
			Name:     "test",
			Quantity: 1,
			Price:    100,
			SubItems: nil,
		},
	}

	output := ConvertOrderItemsToDto(input)

	assert.Equal(t, output[0].Plu, expected[0].Plu)
	assert.Equal(t, output[0].Name, expected[0].Name)
	assert.Equal(t, output[0].Quantity, expected[0].Quantity)
	assert.Equal(t, output[0].Price, expected[0].Price)
}

/*
func TestTwoSimpleItems(t *testing.T) {
	input := []models.OrderItem{
		{
			Id:       1,
			Plu:      "1",
			Name:     "1",
			Price:    100,
			Quantity: 1,
			OrderId:  1,
			ParentId: 0,
		},
		{
			Id:       2,
			Plu:      "2",
			Name:     "2",
			Price:    200,
			Quantity: 2,
			OrderId:  1,
			ParentId: 0,
		},
	}

	expected := []dtos.OrderItem{
		{
			Plu:      "1",
			Name:     "1",
			Quantity: 1,
			Price:    100,
			SubItems: nil,
		},
		{
			Plu:      "2",
			Name:     "2",
			Quantity: 2,
			Price:    200,
			SubItems: nil,
		},
	}

	output := ConvertToOrderItemsDto(input)

	assert.Equal(t, output[0].Plu, expected[0].Plu)
	assert.Equal(t, output[0].Name, expected[0].Name)
	assert.Equal(t, output[0].Quantity, expected[0].Quantity)
	assert.Equal(t, output[0].Price, expected[0].Price)

	assert.Equal(t, output[1].Plu, expected[1].Plu)
	assert.Equal(t, output[1].Name, expected[1].Name)
	assert.Equal(t, output[1].Quantity, expected[1].Quantity)
	assert.Equal(t, output[1].Price, expected[1].Price)
}

func TestNested(t *testing.T) {
	input := []models.OrderItem{
		{
			Id:       1,
			Plu:      "1",
			Name:     "1",
			Price:    100,
			Quantity: 1,
			OrderId:  1,
			ParentId: 0,
		},
		{
			Id:       2,
			Plu:      "2",
			Name:     "2",
			Price:    200,
			Quantity: 2,
			OrderId:  1,
			ParentId: 1,
		},
	}

	expected := []dtos.OrderItem{
		{
			Plu:      "1",
			Name:     "1",
			Quantity: 1,
			Price:    100,
			SubItems: []dtos.OrderItem{
				{
					Plu:      "2",
					Name:     "2",
					Quantity: 2,
					Price:    200,
					SubItems: nil,
				},
			},
		},
	}

	output := ConvertToOrderItemsDto(input)

	assert.Equal(t, output[0].Plu, expected[0].Plu)
	assert.Equal(t, output[0].Name, expected[0].Name)
	assert.Equal(t, output[0].Quantity, expected[0].Quantity)
	assert.Equal(t, output[0].Price, expected[0].Price)

	assert.NotNil(t, output[0].SubItems)
}
*/
