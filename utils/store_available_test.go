package utils

import (
	"testing"
	"time"

	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

type storeAvailableTest struct {
	openingHours []dtos.GetStoreOpeningHour
	day          time.Weekday
	timestamp    int
	available    bool
}

var (
	mockOpeningHoursAlwaysOpen = []storeAvailableTest{
		// Same day matching day and timestamp
		{
			openingHours: []dtos.GetStoreOpeningHour{
				{
					DayOfWeek:      time.Wednesday,
					StartTimestamp: 1000,
					EndTimestamp:   1300,
				},
			},
			day:       time.Wednesday,
			timestamp: 1200,
			available: true,
		},
		// Multiple Openinghours and second one matching
		{
			openingHours: []dtos.GetStoreOpeningHour{
				{
					DayOfWeek:      time.Friday,
					StartTimestamp: 0,
					EndTimestamp:   400,
				},
				{
					DayOfWeek:      time.Sunday,
					StartTimestamp: 1000,
					EndTimestamp:   1300,
				},
			},
			day:       time.Friday,
			timestamp: 300,
			available: true,
		},
		// Multiple Openinghours but not matching
		{
			openingHours: []dtos.GetStoreOpeningHour{
				{
					DayOfWeek:      time.Tuesday,
					StartTimestamp: 1000,
					EndTimestamp:   1159,
				},
				{
					DayOfWeek:      time.Monday,
					StartTimestamp: 100,
					EndTimestamp:   200,
				},
			},
			day:       time.Tuesday,
			timestamp: 1200,
			available: false,
		},
		// One to one not matching
		{
			openingHours: []dtos.GetStoreOpeningHour{
				{
					DayOfWeek:      time.Thursday,
					StartTimestamp: 1000,
					EndTimestamp:   1300,
				},
			},
			day:       time.Saturday,
			timestamp: 1200,
			available: false,
		},
	}
)

func TestStoreAvailable(t *testing.T) {
	for ind, testCase := range mockOpeningHoursAlwaysOpen {
		output := IsStoreAvailable(testCase.openingHours, testCase.day, testCase.timestamp)

		if output != testCase.available {
			t.Errorf("[%v] Expected %v but got %v", ind, testCase.available, output)
		}
	}
}
