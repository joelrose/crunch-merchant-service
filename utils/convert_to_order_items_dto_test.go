package utils

import (
	"database/sql"
	"encoding/json"
	"testing"

	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/stretchr/testify/assert"
)

type ConvertTest struct {
	input    []models.OrderItem
	expected []dtos.OrderItem
}

func TestConvertOrderItemsToDto(t *testing.T) {
	var parent11 sql.NullInt32
	parent11.Int32 = int32(11)
	parent11.Valid = true

	var parent12 sql.NullInt32
	parent12.Int32 = int32(12)
	parent12.Valid = true

	testCases := []ConvertTest{
		{
			input:    []models.OrderItem{},
			expected: []dtos.OrderItem{},
		},
		{
			input: []models.OrderItem{
				{
					Id:       1,
					Plu:      "1",
					Name:     "1",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
			},
			expected: []dtos.OrderItem{
				{
					Id:       1,
					Plu:      "1",
					Name:     "1",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{},
				},
			},
		},
		{
			input: []models.OrderItem{
				{
					Id:       10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
				{
					Id:       11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
				{
					Id:       12,
					Plu:      "12",
					Name:     "12",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
					ParentId: parent11,
				},
				{
					Id:       13,
					Plu:      "13",
					Name:     "13",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
					ParentId: parent12,
				},
			},
			expected: []dtos.OrderItem{
				{
					Id:       10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{},
				},
				{
					Id:       11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{{
						Id:       12,
						Plu:      "12",
						Name:     "12",
						Quantity: 1,
						Price:    100,
						SubItems: []dtos.OrderItem{{
							Id:       13,
							Plu:      "13",
							Name:     "13",
							Quantity: 1,
							Price:    100,
							SubItems: []dtos.OrderItem{},
						},
						},
					},
					},
				},
			},
		},
	}

	for _, testCase := range testCases {
		output := ConvertOrderItemsToDto(testCase.input)

		expectedJson, _ := json.Marshal(testCase.expected)
		outputJson, _ := json.Marshal(output)

		assert.Equal(t, expectedJson, outputJson)
	}
}
