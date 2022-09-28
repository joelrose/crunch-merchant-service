package utils

import (
	"strconv"
	"strings"
	"time"
)

func ConvertToTimestamp(hour int, minute int) int {
	return (hour * 60) + minute
}

func GetPickupTime(averagePickupTime int, timezoneLocale *time.Location) ([]byte, error) {
	return time.Now().In(timezoneLocale).Add(time.Minute * time.Duration(averagePickupTime)).MarshalText()
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

func GetDayAndTimestamp() (time.Weekday, int) {
	berlinTime, err := time.LoadLocation("Europe/Berlin")
	if err != nil {
		panic(err)
	}

	time := time.Now().In(berlinTime)

	timestamp := ConvertToTimestamp(time.Hour(), time.Minute())

	return time.Weekday(), timestamp
}
