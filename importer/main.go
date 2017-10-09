package importer

import (
	"bytes"
	"fmt"
	"os"
	"regexp"

	"github.com/bernardolm/go-csv-tag"
)

func Do(filePath string, build bool) []Workday {
	workdays := []Workday{}

	err := csvtag.Load(csvtag.Config{
		Path:      filePath,
		Dest:      &workdays,
		Separator: ';',
		Modifier:  modifyMyFile,
	})
	if err != nil {
		fmt.Println(err)
	}

	if build {
		for i, v := range workdays {
			fmt.Printf("Workday %d %#v\n", i, v)
		}
	}

	return workdays
}

func modifyMyFile(file *os.File) *string {
	if file == nil {
		return nil
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	contents := buf.String()

	reg, err := regexp.Compile(";\n")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "\n")

	reg, err = regexp.Compile("Saldo.*")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "")

	reg, err = regexp.Compile("\t")
	if err != nil {
		return nil
	}
	contents = reg.ReplaceAllString(contents, "")

	return &contents
}
