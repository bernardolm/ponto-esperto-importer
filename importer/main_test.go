package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	actual := Do("../timesheet.csv", false)

	assert.NotNil(t, actual)
	assert.Len(t, actual, 116)

	assert.Equal(t, "03/08/15", actual[0].Date)
	assert.Equal(t, "-00:25", actual[23].Balance)
	assert.Equal(t, "21:37", actual[34].ExtraOut)
}
