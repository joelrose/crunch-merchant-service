package utils

import "time"

func GetDayAndTimestamp() (time.Weekday, int) {
	time := time.Now()

	timestamp := ConvertToTimestamp(time.Hour(), time.Minute())

	return time.Weekday(), timestamp
}
