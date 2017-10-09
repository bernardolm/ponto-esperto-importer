package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnknownOutputFile(t *testing.T) {
	var file string
	flag.StringVar(&file, "file", "", "")
	flag.Parse()

	assert.Panics(t, func() { main() })
}

func TestUnknownDriver(t *testing.T) {
	var driver string
	flag.StringVar(&driver, "driver", "", "")
	flag.Parse()

	assert.Panics(t, func() { main() })
}

func TestAnotherUnknownDriver(t *testing.T) {
	var driver string
	flag.StringVar(&driver, "foo", "", "")
	flag.Parse()

	assert.Panics(t, func() { main() })
}
