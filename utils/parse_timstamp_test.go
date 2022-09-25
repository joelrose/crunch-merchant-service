package utils

import "testing"

type parseTimestampTest struct {
	input  string
	output int
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
			output: 0,
		},
		{
			input:  "25:01",
			output: 0,
		},
		{
			input:  "-1:-22",
			output: 0,
		},
		{
			input:  "-1:22",
			output: 0,
		},
		{
			input:  "-1:-22",
			output: 0,
		},
		{
			input:  "0:-22",
			output: 0,
		},
		{
			input:  "0:100",
			output: 0,
		},
		// end endge cases
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
