package utils

import (
	"testing"

	"github.com/google/uuid"
	"github.com/joelrose/crunch-merchant-service/models/dtos"
	"github.com/stretchr/testify/assert"
)

type calculateTest struct {
	input    []dtos.OrderItem
	expected int
}

func TestCalculateOrderPrice(t *testing.T) {
	testCases := []calculateTest{
		{
			input:    []dtos.OrderItem{},
			expected: 0,
		},
		{
			input: []dtos.OrderItem{
				{
					Id:       uuid.New(),
					Plu:      "10",
					Name:     "10",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{},
				},
				{
					Id:       uuid.New(),
					Plu:      "11",
					Name:     "11",
					Quantity: 1,
					Price:    100,
					SubItems: []dtos.OrderItem{{
						Id:       uuid.New(),
						Plu:      "12",
						Name:     "12",
						Quantity: 1,
						Price:    100,
						SubItems: []dtos.OrderItem{{
							Id:       uuid.New(),
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
			expected: 400,
		},
	}

	for _, testCase := range testCases {
		actual := CalculateOrderPrice(testCase.input)

		assert.Equal(t, testCase.expected, actual)
	}
}
