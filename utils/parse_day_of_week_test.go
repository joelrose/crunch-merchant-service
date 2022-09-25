package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type dayOfWeekTest struct {
	input    int
	expected int
}

var (
	dayOfWeekTestCases = []dayOfWeekTest{
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
)

func TestParseDeliverectDayOfWeek(t *testing.T) {
	for _, testCase := range dayOfWeekTestCases {
		output := ParseDeliverectDayOfWeek(testCase.input)

		assert.Equal(t, testCase.expected, int(output))
	}
}
