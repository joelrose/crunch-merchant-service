package utils

import "time"

func ConvertToTimestamp(hour int, minute int) int {
	return (hour * 60) + minute
}

func GetDayAndTimestamp() (time.Weekday, int) {
	time := time.Now()

	timestamp := ConvertToTimestamp(time.Hour(), time.Minute())

	return time.Weekday(), timestamp
}
