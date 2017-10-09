package importer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDo(t *testing.T) {
	actual := Do("../file.csv")

	assert.NotNil(t, actual)
	assert.Len(t, actual, 3)
}
