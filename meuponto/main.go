package meuponto

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/bernardolm/ponto-esperto-importer/importer"
	"github.com/bernardolm/ponto-esperto-importer/lib"
)

func Do(workdays []importer.Workday) []Entry {
	entries := []Entry{}

	for _, v := range workdays {
		entries = addEntry(v.In, v, entries)
		entries = addEntry(v.BreakOut, v, entries)
		entries = addEntry(v.BreakIn, v, entries)
		entries = addEntry(v.Out, v, entries)
		entries = addEntry(v.ExtraIn, v, entries)
		entries = addEntry(v.ExtraOut, v, entries)
	}

	if lib.Config.Debug {
		for i, v := range entries {
			fmt.Printf("Entry %d %+v\n", i, v)
		}
	}

	if lib.Config.HasFilePathResult() {
		generateFile(entries)
	}

	return entries
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

func generateFile(entries []Entry) {
	entriesJSON, err := json.Marshal(entries)
	if err != nil {
		panic("can't parse entries to json")
	}

	err = ioutil.WriteFile(lib.Config.FilePathResult(), entriesJSON, 0644)
	if err != nil {
		panic("can't create file")
	}
}

func addEntry(hourMinute string, workday importer.Workday, entries []Entry) []Entry {
	if dateTime := mergeDateTime(workday.Date, hourMinute); dateTime != nil {
		entries = append(entries, Entry{
			Time:    *dateTime,
			Comment: fmt.Sprintf("Using %+v from file %s", workday, lib.Config.FilePath),
		})
	}
	return entries
}
