package utils

import (
	"testing"

	"github.com/joelrose/crunch-merchant-service/dtos"
	"github.com/stretchr/testify/assert"
)

type CalculateTest struct {
	input    []dtos.OrderItem
	expected int
}

func TestCalculateOrderPrice(t *testing.T) {
	testCases := []CalculateTest{
		{
			input:    []dtos.OrderItem{},
			expected: 0,
		},
		{
			input: []dtos.OrderItem{
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
			expected: 400,
		},
	}

	for _, testCase := range testCases {
		actual := CalculateOrderPrice(testCase.input)

		assert.Equal(t, testCase.expected, actual)
	}
}
