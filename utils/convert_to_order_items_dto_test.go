package utils

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/db/models"
	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/stretchr/testify/assert"
)

type ConvertTest struct {
	input    []models.OrderItem
	expected []dtos.OrderItem
}

func TestConvertOrderItemsToDto(t *testing.T) {
	item1, item10, item11, item12, item13 := uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New()

	parent11, parent12 := uuid.NullUUID{}, uuid.NullUUID{}
	parent11.UUID = item11
	parent11.Valid = true

	parent12.UUID = item12
	parent12.Valid = true

	testCases := []ConvertTest{
		{
			input:    []models.OrderItem{},
			expected: []dtos.OrderItem{},
		},
		{
			input: []models.OrderItem{
				{
					Id:       item1,
					Plu:      "1",
					Name:     "1",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
			},
			expected: []dtos.OrderItem{
				{
					Id:       item1,
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
					Id:       item10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
				{
					Id:       item11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
				},
				{
					Id:       item12,
					Plu:      "12",
					Name:     "12",
					Quantity: 1,
					Price:    100,
					OrderId:  1,
					ParentId: parent11,
				},
				{
					Id:       item13,
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
					Id:       item10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{},
				},
				{
					Id:       item11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{{
						Id:       item12,
						Plu:      "12",
						Name:     "12",
						Quantity: 1,
						Price:    100,
						SubItems: []dtos.OrderItem{{
							Id:       item13,
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
