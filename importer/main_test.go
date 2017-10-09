package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	actual := Do("../timesheet.csv")

	assert.NotNil(t, actual)
	assert.Len(t, actual, 116)

	assert.Equal(t, "03/08/15", actual[0].Date)
}
