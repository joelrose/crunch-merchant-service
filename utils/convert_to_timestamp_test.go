package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TimestampTest struct {
	inputHour   int
	inputMinute int
	expected    int
}

func TestConvertToTimestamp(t *testing.T) {
	testCases := []TimestampTest{
		{
			inputHour:   12,
			inputMinute: 30,
			expected:    750,
		},
		{
			inputHour:   0,
			inputMinute: 0,
			expected:    0,
		},
		{
			inputHour:   0,
			inputMinute: 30,
			expected:    30,
		},
		{
			inputHour:   23,
			inputMinute: 0,
			expected:    1380,
		},
	}

	for _, testCase := range testCases {
		actual := ConvertToTimestamp(testCase.inputHour, testCase.inputMinute)

		assert.Equal(t, testCase.expected, actual)
	}
}
