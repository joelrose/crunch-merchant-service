package utils

import (
	"strconv"
	"strings"
	"time"
)

const (
	DeliverectTimeFormat = "2006-01-02T15:04:05Z"
)

func ConvertToTimestamp(hour int, minute int) int {
	return (hour * 60) + minute
}

func GetPickupTime(averagePickupTime int, timezoneLocale *time.Location) time.Time {
	return time.Now().In(timezoneLocale).Add(time.Minute * time.Duration(averagePickupTime))
}

func ParseDeliverectDayOfWeek(day int) time.Weekday {
	if day == 7 {
		return 0
	}

	return time.Weekday(day)
}

func ParseTimestamp(time string) int {
	splitTime := strings.Split(time, ":")

	hour, hErr := strconv.Atoi(splitTime[0])
	minute, mErr := strconv.Atoi(splitTime[1])

	timeConvertError := hErr != nil || mErr != nil
	minuteInvalid := minute > 59 || minute < 0
	hourInvalid := hour > 24 || hour < 0
	timeUpwardsInvalid := hour == 24 && minute > 0

	if timeConvertError || timeUpwardsInvalid || hourInvalid || minuteInvalid {
		return -1
	}

	return ConvertToTimestamp(hour, minute)
}

func GetDayAndTimestamp(timezone *time.Location) (time.Weekday, int) {
	time := time.Now().In(timezone)

	timestamp := ConvertToTimestamp(time.Hour(), time.Minute())

	return time.Weekday(), timestamp
}
