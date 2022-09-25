package utils

import (
	"time"

	"github.com/joelrose/crunch-merchant-service/models/dtos"
)

func IsStoreAvailable(openingHours []dtos.GetStoreOpeningHour, day time.Weekday, timestamp int) bool {
	for _, openingHour := range openingHours {
		if openingHour.DayOfWeek == day {
			if openingHour.StartTimestamp < timestamp && timestamp < openingHour.EndTimestamp {
				return true
			}
		}
	}

	return false
}
