package main

import (
	"flag"
	"fmt"

	"github.com/bernardolm/ponto-esperto-importer/importer"
	"github.com/bernardolm/ponto-esperto-importer/lib"
	"github.com/bernardolm/ponto-esperto-importer/meuponto"
)

func main() {
	fmt.Println("Ponto Esperto importer")

	var file string
	flag.StringVar(&file, "file", "timesheet.csv", "file path with timesheet entries exported by Ponto Esperto app")

	var driver string
	flag.StringVar(&driver, "driver", "meuponto", "driver to transform importer output")

	var debug bool
	flag.BoolVar(&debug, "debug", false, "flag to print debug infos")

	flag.Parse()

	lib.Config = lib.Configuration{
		Debug:    debug,
		Driver:   driver,
		FilePath: file,
	}

	fmt.Printf("Trying import %s with driver %s %s\n", file, driver, lib.Config.DebugMessage())

	lib.Config.CheckFile()

	workdays := importer.Do()

	switch driver {
	case "":
		panic("unknown driver")
	case "meuponto":
		meuponto.Do(workdays)
	default:
		panic("unknown driver")
	}
}
