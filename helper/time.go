package helper

import (
	"os"
	"restfulapi/exception"
	"strconv"
	"strings"
	"time"
)

func StartOfDay(now *time.Time) *time.Time {
	year, month, day := now.Date()
	start := time.Date(year, month, day, 0, 0, 0, 0, now.Location())
	return &start
}

func EndOfDay(now *time.Time) *time.Time {
	year, month, day := now.Date()
	end := time.Date(year, month, day, 23, 59, 59, 999999999, now.Location())
	return &end
}

// ref: https://gemini.google.com/app/15f1e3fadd4c4e3b
func StartOfWeek(now *time.Time) *time.Time {
	// Adjust to Monday.
	weekday := now.Weekday()
	if weekday == time.Sunday {
		weekday = 7 // Adjust Sunday to 7 for easier calculation
	}
	daysToSubtract := int(weekday - time.Monday)
	tm := now.AddDate(0, 0, -daysToSubtract)
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), 0, 0, 0, 0, now.Location())
	return &tm
}

// ref: https://gemini.google.com/app/15f1e3fadd4c4e3b
func EndOfWeek(now *time.Time) *time.Time {
	// Adjust to Sunday.
	weekday := now.Weekday()
	daysToAdd := int(time.Sunday - weekday)
	if daysToAdd < 0 { //handle the case where the current day is after Sunday.
		daysToAdd = daysToAdd + 7
	}
	tm := now.AddDate(0, 0, daysToAdd)
	tm = time.Date(tm.Year(), tm.Month(), tm.Day(), 23, 59, 59, 999999999, now.Location())
	return &tm
}

func StartOfMonth(now *time.Time) *time.Time {
	tm := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())
	return &tm
}

func EndOfMonth(now *time.Time) *time.Time {
	tm := time.Date(now.Year(), now.Month()+1, 0, 23, 59, 59, 999999999, now.Location())
	return &tm
}

func StartOfYear(now *time.Time) *time.Time {
	start := time.Date(now.Year(), 1, 1, 0, 0, 0, 0, now.Location())
	return &start
}

func EndOfYear(now *time.Time) *time.Time {
	end := time.Date(now.Year(), 13, 0, 23, 59, 59, 999999999, now.Location())
	return &end
}

func StrDateToTime(date string) *time.Time {
	var (
		ymdString []string
		ymd       []int = make([]int, 0, 3)
		ymdInt    int
		err       error
		loc       *time.Location
	)
	ymdString = strings.Split(date, "-")
	if len(ymdString) != 3 {
		panic(exception.NewBadRequestException("invalid date format"))
	}
	for i := range ymdString {
		ymdInt, err = strconv.Atoi(ymdString[i])
		if err != nil {
			panic(exception.NewBadRequestException(err.Error()))
		}
		ymd = append(ymd, ymdInt)
	}
	loc, err = time.LoadLocation(os.Getenv("TZ"))
	if err != nil {
		panic(exception.NewBadRequestException(err.Error()))
	}
	tm := time.Date(ymd[0], time.Month(ymd[1]), ymd[2], 0, 0, 0, 0, loc)
	return &tm
}
