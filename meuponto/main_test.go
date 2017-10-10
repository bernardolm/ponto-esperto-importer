package meuponto

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bernardolm/ponto-esperto-importer/importer"
	"github.com/bernardolm/ponto-esperto-importer/lib"
)

func TestParseDate(t *testing.T) {
	actual := parseDate("03/08/999")
	assert.Nil(t, actual)

	actual = parseDate("31/13/99")
	assert.Nil(t, actual)

	actual = parseDate("32/12/99")
	assert.Nil(t, actual)

	actual = parseDate("03/08/15")
	assert.NotNil(t, actual)
	assert.Equal(t, 2015, actual.Year())
	assert.Equal(t, 8, int(actual.Month()))
	assert.Equal(t, 3, actual.Day())

	actual = parseDate("31/12/99")
	assert.NotNil(t, actual)
	assert.Equal(t, 1999, actual.Year())
	assert.Equal(t, 12, int(actual.Month()))
	assert.Equal(t, 31, actual.Day())
}

func TestParseTime(t *testing.T) {
	actual := parseTime("25:00")
	assert.Nil(t, actual)

	actual = parseTime("00:60")
	assert.Nil(t, actual)

	actual = parseTime("00:00")
	assert.NotNil(t, actual)
	assert.Equal(t, 0, actual.Hour())
	assert.Equal(t, 0, actual.Minute())

	actual = parseTime("23:59")
	assert.NotNil(t, actual)
	assert.Equal(t, 23, actual.Hour())
	assert.Equal(t, 59, actual.Minute())

	actual = parseTime("12:12")
	assert.NotNil(t, actual)
	assert.Equal(t, 12, actual.Hour())
	assert.Equal(t, 12, actual.Minute())
}

func TestMergeDateTime(t *testing.T) {
	actual := mergeDateTime("", "")
	assert.Nil(t, actual)

	actual = mergeDateTime("31/12/99", "")
	assert.Nil(t, actual)

	actual = mergeDateTime("", "23:59")
	assert.Nil(t, actual)

	actual = mergeDateTime("31/12/99", "23:59")
	assert.NotNil(t, actual)
	assert.Equal(t, 1999, actual.Year())
	assert.Equal(t, 12, int(actual.Month()))
	assert.Equal(t, 31, actual.Day())
	assert.Equal(t, 23, actual.Hour())
	assert.Equal(t, 59, actual.Minute())

	actual = mergeDateTime("01/01/00", "00:00")
	assert.NotNil(t, actual)
	assert.Equal(t, 2000, actual.Year())
	assert.Equal(t, 1, int(actual.Month()))
	assert.Equal(t, 1, actual.Day())
	assert.Equal(t, 0, actual.Hour())
	assert.Equal(t, 0, actual.Minute())
}

func TestDo(t *testing.T) {
	lib.Config.FilePath = "../timesheet.csv"

	workdays := importer.Do()
	assert.Len(t, workdays, 116)

	lib.Config.FilePath = ""

	actual := Do(workdays)
	assert.Len(t, actual, 502)
}
