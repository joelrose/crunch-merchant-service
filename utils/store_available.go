package utils

import "github.com/joelrose/crunch-merchant-service/models/dtos"

func IsStoreAvailable(openingHours []dtos.GetStoreOpeningHour) bool {
	day, timestamp := GetDayAndTimestamp()
	for _, openingHour := range openingHours {
		if openingHour.DayOfWeek == day {
			continue
		}

		if openingHour.StartTimestamp < timestamp && timestamp < openingHour.EndTimestamp {
			return true
		}
	}

	return false
}
