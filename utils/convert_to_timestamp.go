package utils

import "time"

func ConvertToTimestamp(hour int, minute int) int {
	return (hour * 60) + minute
}

func GetDayAndTimestamp() (int, int) {
	time := time.Now()

	timestamp := ConvertToTimestamp(time.Hour(), time.Minute())

	return int(time.Weekday()), timestamp
}
