package main

import (
	"flag"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	var file string
	flag.StringVar(&file, "file", "", "")
	flag.Parse()
}

func Test(t *testing.T) {
	assert.Panics(t, func() { main() })
}
