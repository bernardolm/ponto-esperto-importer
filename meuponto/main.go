package meuponto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bernardolm/ponto-esperto-importer/importer"
)

func Do(workdays []importer.Workday, filePath string, debug bool) {
	entries := []Entry{}

	for _, v := range workdays {
		if dateTime := mergeDateTime(v.Date, v.In); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}

		if dateTime := mergeDateTime(v.Date, v.BreakOut); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}

		if dateTime := mergeDateTime(v.Date, v.BreakIn); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}

		if dateTime := mergeDateTime(v.Date, v.Out); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}

		if dateTime := mergeDateTime(v.Date, v.ExtraIn); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}

		if dateTime := mergeDateTime(v.Date, v.ExtraOut); dateTime != nil {
			entries = append(entries, Entry{
				Time:    *dateTime,
				Comment: fmt.Sprintf("Using %+v", v),
			})
		}
	}

	if debug {
		for i, v := range entries {
			fmt.Printf("Entry %d %+v\n", i, v)
		}
	}

	if len(filePath) > 0 {
		generateFile(filePath, entries)
	}
}

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

func generateFile(filePath string, entries []Entry) {
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		panic("can't parse entries to json")
	}

	err = ioutil.WriteFile(filePath, entriesJSON, 0644)
	if err != nil {
		panic("can't create file")
	}
}
