package meuponto

import (
	"fmt"
	"time"

	"github.com/bernardolm/ponto-esperto-importer/importer"
)

func Do(workdays []importer.Workday, filePath string, debug bool) {}

func parseDate(date string) *time.Time {
	d, err := time.Parse("02/01/06", date)
	if err != nil {
		fmt.Printf("\n\n err %+v \n\n", err)
		return nil
	}
	return &d
}

func parseTime(hourMinute string) *time.Time {
	t, err := time.Parse("15:04", hourMinute)
	if err != nil {
		fmt.Printf("\n\n err %+v \n\n", err)
		return nil
	}
	return &t
}

func mergeDateTime(dateString string, timeString string) *time.Time {
	if len(dateString) == 0 {
		return nil
	}

	if len(timeString) == 0 {
		return nil
	}

	date := parseDate(dateString)
	if date == nil {
		return nil
	}

	hourMinute := parseTime(timeString)
	if hourMinute == nil {
		return nil
	}

	y, m, d := date.Date()
	result := time.Date(y, m, d, hourMinute.Hour(), hourMinute.Minute(), 0, 0, time.UTC)

	return &result
}
