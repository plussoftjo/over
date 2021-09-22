// Package vendors ..
package vendors

import (
	"strconv"
	"strings"
	"time"
)

// StartOfDay => // Return the Start of The Day
func StartOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC)
}

// EndOfDay => // Return The End Of Day
func EndOfDay(t time.Time) time.Time {
	year, month, day := t.Date()
	return time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenToday => // Return StartOfToday, EndOfToday
func BetwenToday() (time.Time, time.Time) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenYesterDay => // Return StartOfToday, EndOfToday
func BetwenYesterDay() (time.Time, time.Time) {
	currentTime := time.Now()
	year, month, day := currentTime.Date()
	day = day - 1
	return time.Date(year, month, day, 0, 0, 0, 0, time.UTC), time.Date(year, month, day, 23, 59, 59, 59, time.UTC)
}

// BetwenLastSevenDay ..
func BetwenLastSevenDay() (time.Time, time.Time) {
	now := time.Now()

	year, month, day := now.Date()

	startOfLastSevenDay := time.Date(year, month, day-7, 0, 0, 0, 0, time.UTC)
	toToday := time.Date(year, month, day, 23, 59, 59, 59, time.UTC)

	return startOfLastSevenDay, toToday
}

// BetweenThisMonth ..
func BetweenThisMonth() (time.Time, time.Time) {
	now := time.Now()
	currentYear, currentMonth, _ := now.Date()
	currentLocation := now.Location()

	firstOfMonth := time.Date(currentYear, currentMonth, 1, 0, 0, 0, 0, currentLocation)
	lastOfMonth := firstOfMonth.AddDate(0, 1, -1)

	return firstOfMonth, lastOfMonth

}

// BetweenDates ..
func BetweenDates(startDate string, endDate string) (time.Time, time.Time) {
	splitStartDate := strings.Split(startDate, "-")
	splitEndDate := strings.Split(endDate, "-")

	startDateYear, _ := strconv.Atoi(splitStartDate[0])
	startDateMonth, _ := strconv.Atoi(splitStartDate[1])
	startDateDay, _ := strconv.Atoi(splitStartDate[2])

	endDateYear, _ := strconv.Atoi(splitEndDate[0])
	endDateMonth, _ := strconv.Atoi(splitEndDate[1])
	endDateDay, _ := strconv.Atoi(splitEndDate[2])

	isoStartDate := time.Date(startDateYear, time.Month(startDateMonth), startDateDay, 0, 0, 0, 0, time.UTC)
	isoEndDate := time.Date(endDateYear, time.Month(endDateMonth), endDateDay, 23, 59, 59, 59, time.UTC)

	return isoStartDate, isoEndDate
}
