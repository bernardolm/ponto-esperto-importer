package main

import (
	"flag"
	"fmt"

	"github.com/bernardolm/ponto-esperto-importer/importer"
)

func main() {
	fmt.Println("Ponto Esperto importer")

	var file string
	flag.StringVar(&file, "file", "timesheet.csv", "file path with timesheet entries exported by Ponto Esperto app")

	var debug bool
	flag.BoolVar(&debug, "debug", false, "flag to print debug infos")

	flag.Parse()

	importer.Do(file, debug)
}
