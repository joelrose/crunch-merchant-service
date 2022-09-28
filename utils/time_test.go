package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type parseTimestampTest struct {
	input  string
	output int
}

type dayOfWeekTest struct {
	input    int
	expected int
}

type timestampTest struct {
	inputHour   int
	inputMinute int
	expected    int
}

var (
	parseTimestampTestCases = []parseTimestampTest{
		{
			input:  "12:30",
			output: 750,
		},
		{
			input:  "00:00",
			output: 0,
		},
		{
			input:  "00:30",
			output: 30,
		},
		{
			input:  "23:00",
			output: 1380,
		},

		// edge cases
		{
			input:  "24:01",
			output: -1,
		},
		{
			input:  "25:01",
			output: -1,
		},
		{
			input:  "-1:-22",
			output: -1,
		},
		{
			input:  "-1:22",
			output: -1,
		},
		{
			input:  "-1:-22",
			output: -1,
		},
		{
			input:  "0:-22",
			output: -1,
		},
		{
			input:  "0:100",
			output: -1,
		},
		// end edge cases
	}
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

func TestParseTimestamp(t *testing.T) {
	for ind, testCase := range parseTimestampTestCases {
		actual := ParseTimestamp(testCase.input)

		if actual != testCase.output {
			t.Errorf("Test case %d failed. Expected %d, got %d", ind, testCase.output, actual)
		}
	}
}

func TestParseDeliverectDayOfWeek(t *testing.T) {
	for _, testCase := range dayOfWeekTestCases {
		output := ParseDeliverectDayOfWeek(testCase.input)

		assert.Equal(t, testCase.expected, int(output))
	}
}

func TestConvertToTimestamp(t *testing.T) {
	for _, testCase := range timestampTestCases {
		actual := ConvertToTimestamp(testCase.inputHour, testCase.inputMinute)

		assert.Equal(t, testCase.expected, actual)
	}
}
