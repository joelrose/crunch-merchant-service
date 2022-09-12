package utils

func ConvertToTimestamp(hour int, minute int) int {
	return (hour * 60) + minute
}
