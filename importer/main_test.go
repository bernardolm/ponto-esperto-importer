package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/bernardolm/ponto-esperto-importer/lib"
)

func TestDo(t *testing.T) {
	lib.Config.FilePath = "../timesheet.csv"

	actual := Do()

	assert.NotNil(t, actual)
	assert.Len(t, actual, 116)

	assert.Equal(t, "03/08/15", actual[0].Date)
	assert.Equal(t, "-00:25", actual[23].Balance)
	assert.Equal(t, "21:37", actual[34].ExtraOut)
}
