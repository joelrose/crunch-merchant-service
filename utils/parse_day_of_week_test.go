package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type DayOfWeekTest struct {
	input    int
	expected int
}

func TestParseDeliverectDayOfWeek(t *testing.T) {
	testCases := []DayOfWeekTest{
		{
			input:    1,
			expected: 1,
		},
		{
			input:    2,
			expected: 2,
		},
		{
			input:    3,
			expected: 3,
		},
		{
			input:    4,
			expected: 4,
		},
		{
			input:    5,
			expected: 5,
		},
		{
			input:    6,
			expected: 6,
		},
		{
			input:    7,
			expected: 0,
		},
	}

	for _, testCase := range testCases {
		output := ParseDeliverectDayOfWeek(testCase.input)

		assert.Equal(t, testCase.expected, int(output))
	}
}
