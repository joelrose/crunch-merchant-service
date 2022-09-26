package utils

import (
	"strconv"
	"strings"

	"github.com/labstack/gommon/log"
)

func ParseTimestamp(time string) int {
	splitTime := strings.Split(time, ":")

	hour, hErr := strconv.Atoi(splitTime[0])
	minute, mErr := strconv.Atoi(splitTime[1])

	timeConvertError := hErr != nil || mErr != nil
	minuteInvalid := minute > 59 || minute < 0
	hourInvalid := hour > 24 || hour < 0
	timeUpwardsInvalid := hour == 24 && minute > 0

	if timeConvertError || timeUpwardsInvalid || hourInvalid || minuteInvalid {
		log.Errorf("failed to convert time to timestamp: %v", time)
		return 0
	}

	return ConvertToTimestamp(hour, minute)
}
