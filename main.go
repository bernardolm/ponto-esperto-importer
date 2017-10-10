package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/bernardolm/ponto-esperto-importer/importer"
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

	debugMessage := ""
	if debug {
		debugMessage = "in debug mode"
	}

	fmt.Printf("Trying import %s with driver %s %s\n", file, driver, debugMessage)

	workdays := importer.Do(file, debug)

	switch driver {
	case "":
		panic("unknown driver")
	case "meuponto":
		filePathParts := strings.Split(file, ".")
		if len(filePathParts) == 0 {
			panic("unknown output file")
		}
		filePath := fmt.Sprintf("%s_result.json", filePathParts[0])
		meuponto.Do(workdays, filePath, debug)
	default:
		panic("unknown driver")
	}
}
