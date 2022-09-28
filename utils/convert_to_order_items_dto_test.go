package utils

import (
	"encoding/json"
	"testing"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/stretchr/testify/assert"
)

type convertTest struct {
	input    []models.OrderItem
	expected []dtos.OrderItem
}

var (
	mockItemId1, mockItemId10, mockItemId11, mockItemId12, mockItemId13 = uuid.New(), uuid.New(), uuid.New(), uuid.New(), uuid.New()
	mockParentId11, mockParentId12                                      = uuid.NullUUID{UUID: mockItemId11, Valid: true}, uuid.NullUUID{UUID: mockItemId12, Valid: true}
	mockOrderId                                                         = uuid.New()
	convertTestCases                                                    = []convertTest{
		{
			input:    []models.OrderItem{},
			expected: []dtos.OrderItem{},
		},
		{
			input: []models.OrderItem{
				{
					Id:       mockItemId1,
					Plu:      "1",
					Name:     "1",
					Quantity: 1,
					Price:    100,
					OrderId:  mockOrderId,
				},
			},
			expected: []dtos.OrderItem{
				{
					Id:       mockItemId1,
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
					Id:       mockItemId10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					OrderId:  mockOrderId,
				},
				{
					Id:       mockItemId11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					OrderId:  mockOrderId,
				},
				{
					Id:       mockItemId12,
					Plu:      "12",
					Name:     "12",
					Quantity: 1,
					Price:    100,
					OrderId:  mockOrderId,
					ParentId: mockParentId11,
				},
				{
					Id:       mockItemId13,
					Plu:      "13",
					Name:     "13",
					Quantity: 1,
					Price:    100,
					OrderId:  mockOrderId,
					ParentId: mockParentId12,
				},
			},
			expected: []dtos.OrderItem{
				{
					Id:       mockItemId10,
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{},
				},
				{
					Id:       mockItemId11,
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{{
						Id:       mockItemId12,
						Plu:      "12",
						Name:     "12",
						Quantity: 1,
						Price:    100,
						SubItems: []dtos.OrderItem{{
							Id:       mockItemId13,
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
)

func TestConvertOrderItemsToDto(t *testing.T) {
	for _, testCase := range convertTestCases {
		output := ConvertOrderItemsToDto(testCase.input)

		expectedJson, _ := json.Marshal(testCase.expected)
		outputJson, _ := json.Marshal(output)

		assert.Equal(t, expectedJson, outputJson)
	}
}
