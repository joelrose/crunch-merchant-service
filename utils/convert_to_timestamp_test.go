package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type timestampTest struct {
	inputHour   int
	inputMinute int
	expected    int
}

var (
	timestampTestCases = []timestampTest{
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
)

func TestConvertToTimestamp(t *testing.T) {
	for _, testCase := range timestampTestCases {
		actual := ConvertToTimestamp(testCase.inputHour, testCase.inputMinute)

		assert.Equal(t, testCase.expected, actual)
	}
}
