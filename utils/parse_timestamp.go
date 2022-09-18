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

	if hErr != nil || mErr != nil {
		log.Errorf("failed to convert time to timestamp: %v", time)
		return 0
	}

	return ConvertToTimestamp(hour, minute)
}
