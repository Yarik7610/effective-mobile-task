package utils

import "time"

func ParseMonthYearStringToTime(monthYearStr string) (time.Time, error) {
	const mask = "01-2006"
	return time.Parse(mask, monthYearStr)
}
