package utils

import (
	"time"
)

func ParseDeliverectDayOfWeek(day int) time.Weekday {
	if day == 7 {
		return 0
	}

	return time.Weekday(day)
}
